// ============================================================
//  server.js  —  Live Poll WebSocket Server
// ============================================================
//
//  What this file does, in plain English:
//  ----------------------------------------
//  1. Starts an HTTP server (so browsers can load the page)
//  2. Attaches a WebSocket server on top of it
//  3. When a browser connects, it gets added to a "clients" list
//  4. When someone votes or creates a poll, we tell EVERYONE
//     about the update — that's "broadcasting"
//  5. When a browser disconnects, we remove it from the list
//
// ============================================================


// ============================================================
//  server.js  —  Express + WebSocket Live Poll
// ============================================================
//
//  TWO packages. That's it.
//
//    express  →  serves your HTML page (like a waiter bringing the menu)
//    ws       →  handles WebSocket connections (the live phone line)
//
//  How they work together:
//    Express creates an HTTP server.
//    We hand that SAME server to the WebSocket library.
//    So http://localhost:3000  → Express serves the HTML page
//       ws://localhost:3000    → ws handles live connections
//    Both on the same port. No conflict.
//
// ============================================================