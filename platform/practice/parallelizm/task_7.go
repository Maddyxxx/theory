package parallelizm

import (
	//"fmt"
	//"sync"
	"context"
)

//Задача 7


type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

