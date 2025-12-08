# Description

ASCII Art Web is a lightweight yet powerful Go-based web application that transforms plain text into stylized ASCII art. It provides users with a clean, browser-based interface where they can type input text, select from multiple banner styles, and instantly generate ASCII art.

The project is structured to be easily maintainable, scalable, and extensible. It uses Go's efficient I/O handling to serve templates, process text, load banners, and render the final ASCII output. The application is suitable for beginners learning Go web development as well as advanced users looking to integrate ASCII rendering in web services.

Key features include:
- Three banner styles: **Standard**, **Shadow**, and **Thinkertoy**.
- Clean HTML templates for input and output pages.
- Error handling for invalid input, empty text, and unsupported characters.
- Modular code structure: separate files for banner loading, ASCII mapping, routing logic, and utility functions.
- Fast execution thanks to Go’s concurrency-friendly architecture.

# Authors

- ahaddou (handlers and formating)- [GitHub](https://github.com/mrshD3IM05) - [Gitea](https://learn.zone01oujda.ma/git/ahaddou)
- achent (front-end) - [GitHub](https://github.com/chentaymane) - [Gitea](https://learn.zone01oujda.ma/git/achent)
- halhyane (ascii generator) - [GitHub](https://github.com/Houssam-Alhyane) - [Gitea](https://learn.zone01oujda.ma/git/halhyane)

# Usage: How to Run

### Requirements
- Go 1.18 or later
- A modern web browser (Chrome, Firefox, Edge, Safari)
- Git (optional but recommended to clone the repository)

### Installation
1. Clone the repository:
   ```bash
   git clone https://learn.zone01oujda.ma/git/ahaddou/ascii-art-web.git
   cd ascii-art-web
   ```
2. Ensure the required folders exist:
   - `banners/` containing: `standard.txt`, `shadow.txt`, `thinkertoy.txt`
   - `templates/` containing HTML files for input and output pages

3. Build the project:
   ```bash
   go build -o ascii-art-web
   ```

4. Run the server:
   ```bash
   ./ascii-art-web
   ```

### Running Without Building
You can also run directly using:
```bash
go run main.go
```

### Accessing the Web Interface
Once the server is running, open your browser and visit:
```
http://localhost:8080
```

### Usage Flow
1. Enter your text in the textbox.
2. Select a banner style.
3. Press **Generate**.
4. The ASCII version of your input text will be displayed instantly.

# Implementation Details: Algorithm

The ASCII Art Web application operates through a structured pipeline that ensures input text is processed correctly and transformed into ASCII art efficiently.

---
## **1. Request Handling**
- The root route (`/`) serves the main input form using an HTML template.
- The `/ascii-art` route handles POST requests containing:
  - `text`: the user input
  - `banner`: the selected font file
- Input is validated before processing begins.

---
## **2. Input and Validation**
Validation checks include:
- If the text is empty → return an error
- If the text contains tabs or unsupported characters → return an error
- If the text exceeds the character limit → return an error

The algorithm also normalizes Newlines (`\n`) to ensure multi-line ASCII output renders correctly.

---
## **3. Banner File Loading**
Each banner file is a structured text file containing ASCII representations of printable characters.

Banner structure:
- 95 printable characters: ASCII 32 (space) to ASCII 126 (~)
- Each character is represented by **8 lines** of ASCII art
- Characters are stored sequentially in blocks

The algorithm:
1. Opens the selected banner file
2. Splits the file by newline characters (after removing the `\r` to handle exceptions)
3. Groups every 8 lines as the ASCII version of one character
4. Maps each ASCII character to its ASCII art representation

---
## **4. ASCII Rendering Algorithm**
For each line of the input text:
1. For every of the 8 ASCII-art rows:
   - Concatenate the row corresponding to each character
   - Append them horizontally
2. Add newline separators

Example (simplified):
```
Input: "Hi"
Output row 1: H-row1 + I-row1
Output row 2: H-row2 + I-row2
...
```
This preserves proportional spacing and ensures consistent alignment.

---
## **5. Multi-line Support**
If the user enters:
```
Hello
World
```
The algorithm splits the input on `\n` after removing the `\r`, processes each line independently, and joins the final ASCII blocks with a double newline for separation.

---
## **6. HTML Rendering**
The final ASCII output is:
- Escaped to prevent HTML issues
- Embedded inside a `<pre>` tag for monospace formatting
- Passed to the `result.html` template via Go's templating engine

---
## **7. Error Handling and Feedback**
Common errors:
- Missing banner file → server displays an error page
- Invalid characters → message shown on the web UI
- Empty input → friendly error message

These errors are shown through dedicated HTML templates.

---
## **8. Performance Considerations**
- Banner files are kept in memory during request processing for minimal disk I/O
- The application avoids unnecessary allocations
- Concatenation is optimized using Go slices and string builders

---
## **9. Extensibility**
The algorithm is designed to support (in the future):
- New banner styles (just add a `.txt` file)
- New routes or output formats (JSON, file download, etc.)
- API version of the ASCII generator

This modular setup ensures future growth with minimal refactoring.

