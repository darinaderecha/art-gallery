package galleryserver

import (
	artGallery "/home/dar/Desktop/art-gallery/art-gallery/artGallery"
	"sync"
)

type ArtGalleryServer struct {
	picture *Picture
	lock    sync.Mutex
	artGallery.UnimplementedArtGalleryServiceServer
}

type Picture struct {
	pictureId int64
	year      int32
	name      string
	width     float32
	height    float32
}
