package model

type ImageWithComodityID struct {
	ComodityID int
	ImageID    int
	Link       string
}

type temp struct {
	ImagesID   []*int
	Images     []*string
	ComodityID int
}
