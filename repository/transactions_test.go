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

var _ = Describe("TransactionRepository", func() {
	var (
		db    *gorm.DB
		mock  sqlmock.Sqlmock
		sqlDB *sql.DB
		repo  repository.TransactionRepository
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

		repo = repository.NewTransactionRepository(db)
	})

	AfterEach(func() {
		sqlDB.Close()
	})

	Context("AddNewTransaction", func() {
		It("should throw error when inserting same record", func() {
			rows := sqlmock.NewRows([]string{"account_id", "operationtype_id", "amount"}).AddRow(1, 1, 100.0)
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(rows)

			transaction := models.Transaction{
				AccountId:       1,
				OperationTypeId: 1,
				Amount:          100.0,
			}

			_, err := repo.AddNewTransaction(context.TODO(), transaction)
			Expect(err).To(MatchError("record already exists"))
		})

		It("should insert new record", func() {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnError(gorm.ErrRecordNotFound)

			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(
				`INSERT INTO "transactions" ("account_id","amount","operationtype_id") VALUES ($1,$2,$3) RETURNING "transaction_id"`)).
				WithArgs(3, 50.0, 4).
				WillReturnRows(sqlmock.NewRows([]string{"transaction_id"}).AddRow(2))
			mock.ExpectCommit()

			transaction := models.Transaction{
				AccountId:       3,
				OperationTypeId: 4,
				Amount:          50.0,
			}

			result, err := repo.AddNewTransaction(context.TODO(), transaction)
			Expect(err).To(BeNil())
			Expect(result.TransactionId).To(Equal(uint64(2)))
		})
	})
})
