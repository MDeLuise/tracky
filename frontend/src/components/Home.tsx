import { AxiosInstance } from "axios";
import { useEffect, useRef, useState } from "react";
import { Link, NavigateFunction, useNavigate } from "react-router-dom";
import "../style/Home.scss";
import "../style/Base.scss";
import Navbar from "./Navbar";
import RecentTrackers from "./RecentTrackers";
import RecentObservations from "./RecentObservations";
import { Box, Fab, Typography } from "@mui/material";
import AddObservation from "./AddObservation";
import { tracker } from "../interfaces";
import AccountStats from "./AccountStats";
import { isSmallScreen } from "../common";

export default function Home(props: { isLoggedIn: () => boolean, requestor: AxiosInstance; }) {
    let navigate: NavigateFunction = useNavigate();
    const [trackers, setTrackers] = useState<tracker[]>([]);
    const [version, setVersion] = useState<string>("...");
    const [error, setError] = useState<string>();
    const [observationsAdded, setObservationsAdded] = useState<number>(0);
    const windowSize = useRef([window.innerWidth, window.innerHeight]);

    const getCurrentVersion = (): void => {
        props.requestor.get("/info/version")
            .then((response) => {
                setVersion(response.data);
            })
            .catch(error => setError(error));
    };

    const getTrackers = (): void => {
        props.requestor.get("tracker/count")
            .then(response => {
                getAllTrackers(response.data);
            })
            .catch(error => setError(error));
    };

    const getAllTrackers = (total: number): void => {
        props.requestor.get(`tracker?pageSize=${total}&sortBy=lastObservationOn&sortDir=DESC`)
            .then(response => {
                let fetchedTrackers: tracker[] = [];
                response.data.content.forEach((ft: tracker) => {
                    fetchedTrackers.push(ft);
                });
                setTrackers(fetchedTrackers);
            })
            .catch(error => setError(error));
    };

    useEffect(() => {
        if (!props.isLoggedIn()) {
            navigate("/auth");
        } else {
            getCurrentVersion();
            getTrackers();
        }
    }, []);

    if (!props.isLoggedIn()) {
        return <></>;
    }

    return (
        <Box sx={{
            minHeight: "100vh",
        }}>
            <Navbar requestor={props.requestor} mobile={isSmallScreen()} />

            <AddObservation
                requestor={props.requestor}
                trackers={trackers}
                onSuccess={() => {
                    setObservationsAdded(observationsAdded + 1);
                    getTrackers();
                }}
                style={{
                    width: "95% !important",
                }}
            />
            <Box
                id="outer-flex"
                style={{
                    display: 'flex',
                    flexDirection: 'column',
                    padding: '2vw',
                    gap: "10px",
                    justifyContent: "space-between",
                    width: "95%",
                    margin: "0 auto",
                }}
            >
                <Box id="inner-flex" style={{ display: 'flex', width: '100%', justifyContent: "space-between" }}>
                    <Box>
                        <RecentTrackers requestor={props.requestor} trackers={trackers}></RecentTrackers>
                        <RecentObservations requestor={props.requestor} addedObservation={observationsAdded}></RecentObservations>
                    </Box>
                    <Box id="stats">
                        <AccountStats requestor={props.requestor} newData={observationsAdded} />
                    </Box>
                </Box>
                <Box id="credits"
                    sx={{
                        width: '100%',
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                        justifyContent: "end",
                        marginTop: "40px",
                    }}>
                    <Typography variant="body1">
                        Tracky is a project by <Link style={{opacity: ".8", textDecoration: "none", textTransform: "none", color: "inherit"}} to="https://massimilianodeluise.com" target="_blank">MDeLuise</Link>
                    </Typography>
                    <Typography variant="body1">Current version: v{version}</Typography>
                </Box>
            </Box>
        </Box>
    );
}
