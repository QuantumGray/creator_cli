package apirepositories

type ApiRepository struct {
}

var (
	dbURL string = "whatever"
)

type Template struct {
	Identifier string
	Path       string
}

/*
func (ar *ApiRepository) GetTemplate(arg string) (*Template, error) {

	resdb, err := http.Get(dbURL)

	if err != nil {
		return nil, typeserrors.ErrInvalidArgument
	}



	ghURL := fmt.Sprintf(resdb.Body)

	resgh, err := http.Get(ghURL)

	t := &Template{
		"", "",
	}

	return t, nil
}
*/
