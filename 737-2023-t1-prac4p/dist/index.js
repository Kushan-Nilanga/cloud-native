"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
// create express app
const app = (0, express_1.default)();
// parse requests of content-type - application/json
app.use(express_1.default.json());
// get / route
app.get("/", (req, res) => {
    res.send("Hello World! Server is running.");
});
// post /add route
app.post("/add", (req, res) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body;
    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }
    // if everything is fine then send the result
    const response = {
        error: false,
        result: a + b,
    };
    res.send(response);
});
// post /sub route
app.post("/sub", (req, res) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body;
    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }
    // if everything is fine then send the result
    const response = {
        error: false,
        result: a - b,
    };
    res.send(response);
});
// post /mul route
app.post("/mul", (req, res) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body;
    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }
    // if everything is fine then send the result
    const response = {
        error: false,
        result: a * b,
    };
    res.send(response);
});
// post /div route
app.post("/div", (req, res) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body;
    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }
    // if b is 0 then send an error response
    if (b === 0) {
        const response = {
            error: true,
            message: "Cannot divide by 0.",
        };
        return res.send(response);
    }
    // if everything is fine then send the result
    const response = {
        error: false,
        result: a / b,
    };
    res.send(response);
});
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}.`);
});
