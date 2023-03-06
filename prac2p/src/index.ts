
//author: kushan-nilanga

// express imports
import express from "express";

// instantiate express
const app: express.Application = express();

// time logger middleware
app.use((req: express.Request, res: express.Response, next: express.NextFunction) => {
    console.log(`${new Date().toISOString()} [Server Activity]`);
    next();
});

// respond to / with a hello world
app.get("/", (req: express.Request, res: express.Response) => {
    res.send("Hello World!");
});

// start the server
app.listen(3000, () => {
    console.log("Server started on port 3000");
});
