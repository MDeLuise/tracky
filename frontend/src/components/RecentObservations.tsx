import { Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";
import { observation, tracker } from "../interfaces";
import { NavigateFunction, useNavigate } from "react-router-dom";

export default function RecentObservations(props: { requestor: AxiosInstance, addedObservation: number }) {
    const [observations, setObservations] = useState<observation[]>([]);
    const [error, setError] = useState<string>();
    let navigate: NavigateFunction = useNavigate();

    const getRecentObservations = (): void => {
        let fetchedObservations: observation[] = [];
        props.requestor.get("observation?pageSize=5&sortBy=instant&sortDir=DESC")
            .then(response => {
                response.data.content.forEach((ob: observation) => {
                    fetchedObservations.unshift(ob);
                });
                fetchedObservations.reverse();
                setObservations(fetchedObservations);
            })
            .catch(error => setError(error))
    };

    const formatLastObservationInstant = (ob: observation): string => {
        let lastObservationDate = new Date(ob.instant).toLocaleDateString();
        let lastObservationTime = new Date(ob.instant).toLocaleTimeString();
        return `${lastObservationDate}, ${lastObservationTime}`;
    };

    useEffect(() => {
        getRecentObservations();
    }, [props.addedObservation]);

    return (
        <>
            {error && <Typography variant="body1">{"error"}</Typography>}
            <TableContainer component={Paper} sx={{ width: "100%" }}>
                <Table sx={{ minWidth: 650 }} aria-label="last observations" size="small">
                    <caption>Last observations</caption>
                    <TableHead>
                        <TableRow>
                            <TableCell>Value</TableCell>
                            <TableCell>Instant</TableCell>
                            <TableCell>Tracker</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {observations.map((ob: observation) => (
                            <TableRow
                                key={ob.id}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                onClick={() => navigate(`/tracker/${ob.trackerId}`)}
                                className="clickable"
                            >
                                <TableCell>{`${ob.value} ${ob.unit}`}</TableCell>
                                <TableCell>{formatLastObservationInstant(ob)}</TableCell>
                                <TableCell>{ob.trackerName}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}