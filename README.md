[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
<!--   <a href="https://github.com/abhishek-kumaryadav/birthday-messenger">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <h1 align="center">Birthday Messenger</h1>

  <p align="center">
    Automate sending birthday messages from google calendar to whatsapp
    <br />
    <br />
    <a href="#testing">View Demo</a>
    ·
    <a href="https://github.com/abhishek-kumaryadav/birthday-messenger/issues">Report Bug</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#working">Working</a></li>
    <li><a href="#testing">Testing</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
    <li><a href="#what-i-learned">What I Learned</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

![](./src/images/server2FailPass.png)

The project captures all the contacts from Google Calendar who have birthday today and then uses Whatapp API to send them birthday wishes. The project provides a docker implementation for jobs scheduled in golang.

### Built With

- [Go Dev](https://go.dev/doc/)
- [Docker](https://docs.docker.com/)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Install following packages in your environment:

- Install docker

### Installation

1. Clone the repo:
   ```sh
   git clone https://github.com/abhishek-kumaryadav/birthday-messenger.git
   ```
2. Create facebook credentials following the link https://developers.facebook.com/docs/whatsapp/cloud-api/get-started#set-up-developer-assets
    - Add the credentials in birthday-messenger.properties file
3. Create google credentials following the link https://developers.google.com/people/quickstart/go
    - Add the credentials in birthday-messenger.properties file
    - Add the credentials.json
4. Set the job frequency to per day once
<!-- USAGE EXAMPLES -->

## Usage

Run the below command in your command line terminal to launch the program:

```
docker build -t go-docker-birthday-messenger .
docker run go-docker-birthday-messenger
```

<!-- WORKING -->

## Working

- Docker image runs a golang application
- The golang application creates a gocron job which runs once per day
- The job fetches contacts from google api and sends whatsapp message to those contacts
<!-- TESTING -->

## Testing

Testing was done by locally running the docker image

<!-- ROADMAP -->

## Roadmap

See the [open issues](https://github.com/abhishek-kumaryadav/birthday-messenger/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Abhishek Kumar Yadav - [twitter/abhishek-kumaryadav](https://twitter.com/abhishek-kumaryadav) - abhk943@gmail.com

Project Link: [https://github.com/abhishek-kumaryadav/birthday-messenger](https://github.com/abhishek-kumaryadav/birthday-messenger)

<!-- ACKNOWLEDGEMENTS -->

## Acknowledgements

- [Dockerizing Go Application](https://blog.logrocket.com/dockerizing-go-application/)
- [Google Developers](https://developers.google.com/)

<!-- WHAT I LEARNED -->

## What I Learned

- **Golang programming language**
- **Docker utilization for running applications**
  <!-- MARKDOWN LINKS & IMAGES -->
  <!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/abhishek-kumaryadav/birthday-messenger.svg?style=for-the-badge
[contributors-url]: https://github.com/abhishek-kumaryadav/birthday-messenger/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/abhishek-kumaryadav/birthday-messenger.svg?style=for-the-badge
[forks-url]: https://github.com/abhishek-kumaryadav/birthday-messenger/network/members
[stars-shield]: https://img.shields.io/github/stars/abhishek-kumaryadav/birthday-messenger.svg?style=for-the-badge
[stars-url]: https://github.com/abhishek-kumaryadav/birthday-messenger/stargazers
[issues-shield]: https://img.shields.io/github/issues/abhishek-kumaryadav/birthday-messenger.svg?style=for-the-badge
[issues-url]: https://github.com/abhishek-kumaryadav/birthday-messenger/issues
[license-shield]: https://img.shields.io/github/license/abhishek-kumaryadav/birthday-messenger.svg?style=for-the-badge
[license-url]: https://github.com/abhishek-kumaryadav/birthday-messenger/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/abhishek-kumaryadav
