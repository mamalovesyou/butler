import Error from "../components/error";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";


const ErrorPage = () => {

    const [errorCode, setErrorCode] = useState("500")
    const {code} = useParams();

    useEffect(() => setErrorCode(code), [code])
    console.log(code, errorCode)
    switch (errorCode) {
        case "404":
            return <Error
                title={"Page not found"}
                description={"The page you are looking for isn’t here. You either tried some shady route or you came here by mistake.\n" +
                "Whichever it is, try using the navigation"}
            />
        case "403":
            return <Error
                title={"You can't do that"}
                description={"It seems that you are trying to do something you can't."}
            />
        default:
            return <Error />
    }
};

export default ErrorPage;

