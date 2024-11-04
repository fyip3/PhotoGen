# AI Photo Generator

This project generates custom AI-powered profile pictures based on user-provided descriptions. It leverages a Go-based backend to interact with the OpenAI API and a React frontend to deliver an intuitive and interactive user experience.

---

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
---

## Overview

The AI Photo Generator allows users to upload an image, enter a description for the desired style, and receive AI-generated variations of profile pictures. This tool is versatile for users looking to create professional, styled, or artistic headshots for various online profiles.

## Features

- **AI-Generated Photos**: Generates multiple photo variations based on a single description, allowing users to explore a range of professional or artistic headshots.
- **Concurrent Processing**: Utilizes Go’s concurrency model to handle multiple image generation tasks in parallel, optimizing performance and reducing wait times for users.
- **Dynamic User Interface**: An interactive UI with loading indicators, success/failure notifications, and animations for a seamless user experience.
- **Temporary Database Storage**: Stores generated images temporarily in MongoDB, allowing users to retrieve images within a limited session window.
- **Error Handling and User Feedback**: Provides clear notifications and alerts, ensuring smooth interactions and helping users understand the process.


## Tech Stack

- **Backend**: Go, Gorilla Mux, OpenAI API
- **Frontend**: React, Material UI, Axios
- **Database**: MongoDB

