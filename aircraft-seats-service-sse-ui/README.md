# Aircraft Seats Map

A React application for visualizing and interacting with aircraft seat layouts in real time.  
This project uses **React**, **GraphQL**, and **Apollo Client** with subscriptions over **SSE** for live seat status updates.

## Features

- Real-time seat status updates using GraphQL subscriptions (SSE)
- Interactive seat selection and status display
- Modular React components

## Tech Stack

- [React](https://react.dev/)
- [Apollo Client](https://www.apollographql.com/docs/react/)
- [GraphQL](https://graphql.org/)
- [graphql-sse](https://github.com/enisdenjo/graphql-sse) for subscriptions over SSE

## Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) (v18 or newer recommended)
- [npm](https://www.npmjs.com/)

### Installation

1. **Clone the repository:**
   ```sh
   git clone git@github.com:hadeed85/apidays-demo.git
   cd aircraft-seats-service-sse-ui
   ```

2. **Install dependencies:**
   ```sh
   npm install
   ```   

### Running the Application

1. **Start the development server:**
   ```sh
   npm start
   ```   

2. Open your browser and go to [http://localhost:3000](http://localhost:3000).

**Note:**  
This UI expects a GraphQL backend running at `https://graphql-gateway.app.lan:4444/graphql`

## Project Structure

- `src/components/` – React components for seats, rows, and wings
- `src/index.js` – Apollo Client setup with HTTP and WebSocket links

### Build the application

   ```sh
   npm run build
   ```   
