import express, { Request, Response } from "express";
import passport from "passport";
import passportJWT from "passport-jwt";
import jwt from "jsonwebtoken";

// authentication strategy
const JWTStrategy = passportJWT.Strategy;
const ExtractJWT = passportJWT.ExtractJwt;

// interfaces
// response interface
interface IRensponse {
    error: boolean;
    result?: number;
    message?: string;
}

// request body interface for calculation
interface ICalcRequest {
    a: number;
    b: number;
}

// request body interface for authentication
interface IAuthRequest {
    username: string;
    password: string;
}

// passport configuration
const jwtOps = {
    jwtFromRequest: ExtractJWT.fromAuthHeaderAsBearerToken(),
    secretOrKey: "secret",
};

// use the JWT strategy
passport.use(
    new JWTStrategy(jwtOps, (jwtPayload, done) => {
        // find the user in db if needed. This functionality may be omitted if you store everything you'll need in JWT payload.
        return done(null, jwtPayload);
    })
);

// create express app
const app = express();
app.use(express.json());

// get /auth route to authenticate user
app.get("/auth", (req: Request, res: Response) => {
    // this will expect the request body to be of type IAuthRequest
    // if not it will send an response of type IRensponse with error set to true
    const { username, password } = req.body as IAuthRequest;

    // if username or password is not provided then send an error response
    if (!username || !password) {
        const response: IRensponse = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }

    // if username is not admin or password is not admin then send an error response
    // this is where you would check the database for the user
    if (username !== "admin" || password !== "admin") {
        const response: IRensponse = {
            error: true,
            message: "Invalid username or password.",
        };
        return res.send(response);
    }

    // generate a token that is valid for 1 hour
    const token = jwt.sign({ username, password }, "secret", {
        expiresIn: 60,
    });

    // if everything is fine then send the token
    const response: IRensponse = {
        error: false,
        message: token,
    };

    res.send(response);
});

// post /add route
app.post(
    "/add",
    passport.authenticate("jwt", { session: false }),
    (req: Request, res: Response) => {
        // this will expect the request body to be of type ICalcRequest
        // if not it will send an response of type IRensponse with error set to true
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
    }
);

// post /sub route
app.post("/sub", (req: Request, res: Response) => {
    // this will expect the request body to be of type ICalcRequest
    // if not it will send an response of type IRensponse with error set to true
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
app.post(
    "/mul",
    passport.authenticate("jwt", { session: false }),
    (req: Request, res: Response) => {
        // this will expect the request body to be of type ICalcRequest
        // if not it will send an response of type IRensponse with error set to true
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
    }
);

// post /div route
app.post(
    "/div",
    passport.authenticate("jwt", { session: false }),
    (req: Request, res: Response) => {
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
    }
);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}.`);
});
