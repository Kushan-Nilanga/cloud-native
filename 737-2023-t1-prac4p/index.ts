import express, { Request, Response } from "express";

// create express app
const app = express();

// parse requests of content-type - application/json
app.use(express.json());

// get / route
app.get("/", (req: Request, res: Response) => {
    res.send("Hello World! Server is running.");
});

interface IRensponse {
    error: boolean;
    result?: number;
    message?: string;
}

interface ICalcRequest {
    a: number;
    b: number;
}

// post /add route
app.post("/add", (req: Request, res: Response) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body as ICalcRequest;

    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response: IRensponse = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }

    // if everything is fine then send the result
    const response: IRensponse = {
        error: false,
        result: a + b,
    };

    res.send(response);
});

// post /sub route
app.post("/sub", (req: Request, res: Response) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body as ICalcRequest;

    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response: IRensponse = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }

    // if everything is fine then send the result
    const response: IRensponse = {
        error: false,
        result: a - b,
    };

    res.send(response);
});

// post /mul route
app.post("/mul", (req: Request, res: Response) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body as ICalcRequest;

    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response: IRensponse = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }

    // if everything is fine then send the result
    const response: IRensponse = {
        error: false,
        result: a * b,
    };

    res.send(response);
});

// post /div route
app.post("/div", (req: Request, res: Response) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
    // this is done by using the type assertion syntax in typescript
    // the body value types must match the ICalcRequest interface
    const { a, b } = req.body as ICalcRequest;

    // if a or b is not a number then send an error response
    if (isNaN(a) || isNaN(b)) {
        const response: IRensponse = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }

    // if b is 0 then send an error response
    if (b === 0) {
        const response: IRensponse = {
            error: true,
            message: "Cannot divide by 0.",
        };
        return res.send(response);
    }

    // if everything is fine then send the result
    const response: IRensponse = {
        error: false,
        result: a / b,
    };

    res.send(response);
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}.`);
});
