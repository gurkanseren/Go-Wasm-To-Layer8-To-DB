package models

const ImageURLOne = "https://c8.alamy.com/comp/2K3MF4R/300-2006-300-movie-poster-gerard-butler-2K3MF4R.jpg"

const ImageURLTwo = "https://images.moviesanywhere.com/2184f0c69e02ee614a3d9df0d117f675/55e42cb4-4df6-4749-acf9-432094cb239f.jpg"

const ImageURLThree = "https://m.media-amazon.com/images/M/MV5BMTA4MjIyZWEtZjYwMS00ZmQ1LWJiMDEtMWNiNTI5NWE3OGJjXkEyXkFqcGdeQXVyNjk1Njg5NTA@._V1_FMjpg_UX1000_.jpg"

const ImageURLFour = "https://m.media-amazon.com/images/M/MV5BNDIzMTk4NDYtMjg5OS00ZGI0LWJhZDYtMzdmZGY1YWU5ZGNkXkEyXkFqcGdeQXVyMTI5NzUyMTIz._V1_.jpg"

type ImageURLResponse struct {
	URL string `json:"url"`
}

type ECDHKeyExchangeRequest struct {
	PrivKeySlaveD []byte `json:"PrivKeySlaveD"`
}

type ECDHKeyExchangeOutput struct {
	PrivKeyServerD []byte `json:"PrivKeyServerD"`
}
