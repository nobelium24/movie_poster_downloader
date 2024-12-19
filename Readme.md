# Movie Poster Downloader

Movie Poster Downloader is a command-line tool written in Go that fetches and downloads the poster of a specified movie using the OMDB API.

## Requirements

- Go 1.16 or later
- OMDB API key (Sign up at [OMDB API](https://www.omdbapi.com/apikey.aspx) to get your API key)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/MoviePosterDownloader.git
   cd MoviePosterDownloader
   ```

2. Build the tool:
   ```bash
   go build -o movieposterdownloader
   ```

## Usage

### Command Format

```bash
./movieposterdownloader <movie_title>
```

- Replace `<movie_title>` with the title of the movie whose poster you want to download.
- Ensure your OMDB API key is included in the code. Update the `apikey` parameter in the request URL inside the code if needed.

### Example

To fetch the poster for the movie "Inception":

```bash
./movieposterdownloader "Inception"
```

This will download the poster image as `poster.jpg` in the current working directory.

## File Structure

```
MoviePosterDownloader/
├── go.mod         # Module file
├── main.go        # Tool implementation
└── README.md      # Documentation
```

## Known Issues

- Ensure the OMDB API key is valid and has not exceeded its request limit.
- The movie title must be enclosed in quotes if it contains spaces.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributions

Contributions are welcome! Feel free to fork the repository and submit a pull request.

## Contact

For any questions or issues, contact the author at [ogunbaja24@gmail.com].
