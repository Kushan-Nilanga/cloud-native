"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const passport_1 = __importDefault(require("passport"));
const passport_jwt_1 = __importDefault(require("passport-jwt"));
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
// authentication strategy
const JWTStrategy = passport_jwt_1.default.Strategy;
const ExtractJWT = passport_jwt_1.default.ExtractJwt;
const jwtOps = {
    jwtFromRequest: ExtractJWT.fromAuthHeaderAsBearerToken(),
    secretOrKey: "secret",
};
passport_1.default.use(new JWTStrategy(jwtOps, (jwtPayload, done) => {
    // find the user in db if needed. This functionality may be omitted if you store everything you'll need in JWT payload.
    return done(null, jwtPayload);
}));
// create express app
const app = (0, express_1.default)();
app.use(express_1.default.json());
// get / route
app.get("/", (req, res) => {
    res.send("Hello World! Server is running.");
});
// get /auth route to authenticate user
app.get("/auth", (req, res) => {
    // this will expect the request body to be of type IAuthRequest
    // if not it will send an response of type IRensponse with error set to true
    const { username, password } = req.body;
    // if username or password is not provided then send an error response
    if (!username || !password) {
        const response = {
            error: true,
            message: "Invalid request body.",
        };
        return res.send(response);
    }
    // if username is not admin or password is not admin then send an error response
    // this is where you would check the database for the user
    if (username !== "admin" || password !== "admin") {
        const response = {
            error: true,
            message: "Invalid username or password.",
        };
        return res.send(response);
    }
    // generate a token that is valid for 1 hour
    const token = jsonwebtoken_1.default.sign({ username, password }, "secret", {
        expiresIn: 60,
    });
    // if everything is fine then send the token
    const response = {
        error: false,
        message: token,
    };
    res.send(response);
});
// post /add route
app.post("/add", passport_1.default.authenticate("jwt", { session: false }), (req, res) => {
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
app.post("/mul", passport_1.default.authenticate("jwt", { session: false }), (req, res) => {
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
app.post("/div", passport_1.default.authenticate("jwt", { session: false }), (req, res) => {
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
