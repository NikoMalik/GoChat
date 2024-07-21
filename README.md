# GoChat: Real-Time Chat Application

GoChat is a real-time chat application developed using a modern tech stack featuring Go+Templ+HTMX+Alpine.js+TailwindCss, and WebSockets. This project showcases a scalable, interactive chat system with efficient server-side rendering and dynamic front-end updates.



![image](https://github.com/user-attachments/assets/82e993a5-14bc-405d-b49a-f3a6e8d05356)


![image](https://github.com/user-attachments/assets/ae3d87c5-b772-48a2-aa4d-64ee6826679a)


![image](https://github.com/user-attachments/assets/19b682e9-4976-40f3-8aaa-7812f6777784)



![image](https://github.com/user-attachments/assets/f3de35e7-e59c-474d-bdfa-1943476334bc)




## Features

- **Real-Time Messaging**: Leveraging WebSockets for instantaneous, bidirectional communication.
- **Dynamic Content Updates**: Using HTMX to seamlessly update parts of the web page without full reloads.
- **Reactive UI**: Enhanced with Alpine.js for minimal overhead and a smooth user experience.
- **Server-Side Templating**: Implementing Templ to render dynamic content server-side.
- **Responsive Design**: Styled with Tailwind CSS for a modern, responsive interface.
- **Rate Limiting**: Protects against spam with configurable message rate limits.

## Tech Stack

- **Go**: Backend development for a robust and efficient server-side experience.
- **Templ**: A templating engine for Go, simplifying dynamic content rendering.
- **HTMX**: For AJAX requests and dynamic content updates without page reloads.
- **Alpine.js**: A minimal JavaScript framework to add interactivity with ease.
- **WebSockets**: Provides real-time messaging capabilities.
- **Tailwind CSS**: A utility-first CSS framework for designing custom, responsive interfaces.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- A web browser

### Installation

1.  **Clone the Repository:**

    ```bash
    git clone https://github.com/NikoMalik/GoChat
    ```

2.  **Install Dependencies:**

    ```bash
    go install github.com/air-verse/air@latest
    go install github.com/a-h/templ/cmd/templ@latest
    npm install


    ```

3.  **Run the Application:**

    You will need to open three terminals to run different tasks:

    - **Terminal 1: Watch for Templ**

      ```bash
      make templWatch
      ```

    - **Terminal 2: Run the Go server**

      ```bash
      air
      ```

    - **Terminal 3: Build Tailwind CSS**

      ```bash
      make tailwind
      ```

4.  **Open Your Browser:**

    Visit `http://localhost:8000` to access GoChat.

## Usage

- **Chat with Others**: Enter your username and start sending messages in real-time.
- **View Online Users**: See who is online and engage with them directly.
- **Rate Limiting**: Prevents abuse by enforcing message limits.

## Contributing

We welcome contributions to improve GoChat! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b your-feature`).
3. Make your changes.
4. `git add .`
5. Commit your changes (`git commit -m 'Add new feature'`).
6. Push to the branch (`git push origin your-feature`).
7. Open a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
