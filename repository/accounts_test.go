package repository_test

import (
	"context"
	"database/sql"
	"pismo-service/models"
	"pismo-service/repository"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _ = Describe("Accounts", func() {
	var (
		db    *gorm.DB
		mock  sqlmock.Sqlmock
		sqlDB *sql.DB
		repo  repository.AccountRepository
		err   error
	)

	BeforeEach(func() {
		sqlDB, mock, err = sqlmock.New()
		Expect(err).To(BeNil())

		dialector := postgres.New(postgres.Config{
			Conn:                 sqlDB,
			PreferSimpleProtocol: true,
		})

		db, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		Expect(err).To(BeNil())

		repo = repository.NewAccountRepository(db)
	})

	AfterEach(func() {
		sqlDB.Close()
	})

	Context("GetOneAccountById", func() {
		It("should retrieve valid account", func() {
			rows := sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow(1, "abcde12345")
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(rows)

			account := models.Account{
				DocumentNumber: "abcde12345",
				AccountId:      1,
			}

			result, err := repo.GetOneAccountById(context.TODO(), account.AccountId)
			Expect(err).To(BeNil())
			Expect(result.DocumentNumber).To(Equal("abcde12345"))
		})

		It("should return err if query fails", func() {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnError(gorm.ErrRecordNotFound)

			account := models.Account{
				DocumentNumber: "abcde12345",
				AccountId:      1,
			}

			_, err := repo.GetOneAccountById(context.TODO(), account.AccountId)
			Expect(err).NotTo(BeNil())
			Expect(err).To(MatchError(gorm.ErrRecordNotFound))
		})
	})

	Context("CreateNewAccount", func() {
		It("should throw error when inserting same record", func() {
			rows := sqlmock.NewRows([]string{"account_id", "document_number"}).AddRow(1, "abcdef123456")
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(rows)

			transaction := models.Account{
				AccountId:      1,
				DocumentNumber: "abcdef123456",
			}

			_, err := repo.CreateNewAccount(context.TODO(), transaction)
			Expect(err).To(MatchError("record already exists"))
		})

		It("should insert new record", func() {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnError(gorm.ErrRecordNotFound)

			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(
				`INSERT INTO "accounts" ("document_number") VALUES ($1) RETURNING "account_id"`)).
				WithArgs("abdcef654321").
				WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow(2))
			mock.ExpectCommit()

			account := models.Account{
				DocumentNumber: "abdcef654321",
			}

			result, err := repo.CreateNewAccount(context.TODO(), account)
			Expect(err).To(BeNil())
			Expect(result.AccountId).To(Equal(uint64(2)))
		})
	})

})
