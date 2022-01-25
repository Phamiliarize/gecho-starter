package interactor

// type BookInteractor interface {
// 	GetBookInteractor(request *GetBookRequest) *GetBookResponse
// }

type Book struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Read bool   `json:"read"`
}

type GetBookRequest struct {
	ID uint32 `param:"id"`
}

type GetBookResponse struct {
	Status int
	Body   Book
}

func GetBookInteractor(request *GetBookRequest) *GetBookResponse {
	var response GetBookResponse
	return &response
}
