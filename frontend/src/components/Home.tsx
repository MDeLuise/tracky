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
import AddIcon from '@mui/icons-material/Add';

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
        <>
            <Navbar requestor={props.requestor}></Navbar>

            {/* {windowSize.current[0] <= 1357 &&
                <Fab
                    color="primary"
                    aria-label="add entity"
                    onClick={() => navigate("/add")}
                    sx={{
                        position: "fixed",
                        bottom: "20px",
                        right: "20px"
                    }}
                >
                    <AddIcon />
                </Fab>
                ||
                <AddObservation
                    requestor={props.requestor}
                    trackers={trackers}
                    onSuccess={() => {
                        setObservationsAdded(observationsAdded + 1);
                        getTrackers();
                    }}
                    style={{ marginTop: "30px" }}
                />
            } */}
            <AddObservation
                requestor={props.requestor}
                trackers={trackers}
                onSuccess={() => {
                    setObservationsAdded(observationsAdded + 1);
                    getTrackers();
                }}
                style={{ marginTop: "30px" }}
            />
            <Box id="outer-flex" style={{ display: 'flex', flexDirection: 'column', padding: '2vw', gap: "10px", minHeight: "93vh", justifyContent: "space-between" }}>
                <Box id="inner-flex" style={{ display: 'flex', width: '100%', justifyContent: "space-between" }}>
                    <Box>
                        <RecentTrackers requestor={props.requestor} trackers={trackers}></RecentTrackers>
                        <RecentObservations requestor={props.requestor} addedObservation={observationsAdded}></RecentObservations>
                    </Box>
                    <Box id="stats">
                        <AccountStats requestor={props.requestor} newData={observationsAdded} />
                    </Box>
                </Box>
                <Box id="credits" sx={{ width: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: "end" }}>
                    <Typography variant="body1">Tracky is a project by <Link to="https://github.com/MDeLuise/tracky#tracky" target="_blank">MDeLuise</Link></Typography>
                    <Typography variant="body1">Current version: v{version}</Typography>
                </Box>
            </Box>
        </>
    );
}
