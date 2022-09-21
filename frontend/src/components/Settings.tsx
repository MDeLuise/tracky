import { AxiosInstance } from "axios";
import { useEffect } from "react";
import { NavigateFunction, useNavigate } from "react-router-dom";
import "../style/Settings.scss";
import Navbar from "./Navbar";

export default function Settings(props: { isLoggedIn: () => boolean, requestor: AxiosInstance }) {
    let navigate: NavigateFunction = useNavigate();


    useEffect(() => {
        if (!props.isLoggedIn()) {
            navigate("/");
        }

    }, []);

    return (
        <>
            <Navbar requestor={props.requestor}></Navbar>
        </>
    );
}
