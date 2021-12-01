package entity

type TransactionRepository interface {
	Insert(id string, accountId string, amount float64, status string, errorMessage string) error
}

// Generating mock for repository with mockgen
// go install github.com/golang/mock/mockgen@v1.6.0
// mockgen -source=entity/repository.go -destination=entity/mock/mock.go

// in case of error, add binary paths
// export GOPATH="$HOME/go"
// PATH="$GOPATH/bin:$PATH"
