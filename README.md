# LinkedIn Profile Picture Generator

This project generates LinkedIn-style profile pictures using AI based on user-provided descriptions. It leverages a Go-based backend to interact with the OpenAI API, and a React frontend to provide an intuitive, interactive user experience.

---

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

---

## Overview

The LinkedIn Profile Picture Generator allows users to upload an image, input a description of their desired style, and receive several AI-generated profile pictures. It's ideal for users looking to create professional or styled headshots for social media profiles.

## Features

- **AI-Generated Profile Pictures**: Generate multiple profile pictures based on a single description.
- **Dynamic User Experience**: Interactive UI with loading indicators, notifications, and smooth animations.
- **Error Handling and User Feedback**: Displays errors and success notifications for enhanced user experience.
- **Responsive Design**: Optimized for both desktop and mobile devices.

## Tech Stack

- **Backend**: Go, Gorilla Mux, OpenAI API
- **Frontend**: React, Material UI, Axios
- **Database**: MongoDB (optional, for saving images by session)

## Project Structure

```plaintext
linkedin-profile-generator/
├── backend/               # Go server files
│   ├── main.go            # Server setup and CORS handling
│   ├── handlers.go        # HTTP handlers for endpoints
│   ├── image_processor.go # AI image processing logic
│   ├── database.go        # Database connection and models (optional)
│   └── go.mod             # Go module file
├── frontend/              # React application files
│   ├── public/            # Static files
│   └── src/
│       ├── App.js         # Main React app
│       ├── components/    # React components
│       │   ├── ImageUploader.js
│       │   └── ImageGrid.js
│       ├── theme.js       # Custom Material UI theme
│       ├── App.css        # Main CSS styles
│       └── index.js       # React entry point
├── .env                   # Environment variables
└── README.md              # Project documentation
