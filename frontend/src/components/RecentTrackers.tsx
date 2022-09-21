import { Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";
import { tracker } from "../interfaces";
import { NavigateFunction, useNavigate } from "react-router-dom";

export default function RecentTrackers(props: { requestor: AxiosInstance, trackers: tracker[] }) {
    const [trackers, setTrackers] = useState<tracker[]>([]);
    const [error, setError] = useState<string>();
    let navigate: NavigateFunction = useNavigate();

    const getRecentTrackers = (): void => {
        const dateComparator = (d1: Date, d2: Date): -1 | 0 | 1 => {
            if (d1 < d2) {
                return 1;
            } else if (d2 < d1) {
                return -1;
            } else {
                return 0;
            }
        };

        const trackerLastObservationSorting = (tr1: tracker, tr2: tracker): -1 | 0 | 1 => {
            if (tr1.lastObservation == undefined && tr2.lastObservation == undefined) {
                return 0;
            } else if (tr1.lastObservation == undefined) {
                return 1;
            } else if (tr2.lastObservation == undefined) {
                return -1;
            } else {
                return dateComparator(tr1.lastObservation!.instant, tr2.lastObservation!.instant);
            }
        };

        let allTrackers: tracker[] = [];
        props.trackers.forEach(tr => allTrackers.push(tr));
        allTrackers.sort((tr1: tracker, tr2: tracker) => trackerLastObservationSorting(tr1, tr2));
        setTrackers(allTrackers.slice(0, 6));
    };

    const formatLastObservationValue = (tr: tracker): string => {
        if (tr.lastObservation == undefined) {
            return "-";
        }
        return `${tr.lastObservation.value} ${tr.unit}`;
    }

    const formatLastObservationInstant = (tr: tracker): string => {
        if (tr.lastObservation == undefined) {
            return "-";
        }
        let lastObservationDate = new Date(tr.lastObservation!.instant).toLocaleDateString();
        let lastObservationTime = new Date(tr.lastObservation!.instant).toLocaleTimeString();
        return `${lastObservationDate}, ${lastObservationTime}`;
    };

    useEffect(() => {
        getRecentTrackers();
    }, [props.trackers]);

    return (
        <>
            {error && <Typography variant="body1">{"error"}</Typography>}
            <TableContainer component={Paper} sx={{ width: "100%", marginBottom: "20px" }}>
                <Table sx={{ minWidth: 650 }} aria-label="Table of recent trackers" size="small">
                    <caption>Last trackers</caption>
                    <TableHead>
                        <TableRow>
                            <TableCell style={{ width: "fit-content" }}>Name</TableCell>
                            <TableCell style={{ width: "fit-content" }}>Last value</TableCell>
                            <TableCell style={{ width: "fit-content" }}>Updated on</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {trackers.map((tr: tracker) => (
                            <TableRow
                                key={tr.id}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                onClick={() => navigate(`/tracker/${tr.id}`)}
                                className="clickable"
                            >
                                <TableCell style={{ width: "fit-content" }} component="th" scope="row">
                                    {tr.name}
                                </TableCell>
                                <TableCell style={{ width: "fit-content" }}>{formatLastObservationValue(tr)}</TableCell>
                                <TableCell style={{ width: "fit-content" }}>{formatLastObservationInstant(tr)}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer >
        </>
    )
}