import { Typography, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from "@mui/material";
import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";

export default function Statistics(props: { colorMode: any, requestor: AxiosInstance }) {
    // const [statistics, setStatistics] = useState<observation[]>([]);
    // const [error, setError] = useState<string>();

    // const getStatistics = (): void => {

    // };

    // useEffect(() => {
    //     getRecentObservations();
    // }, []);

    // return (
    //     <>
    //         {error && <Typography variant="body1">{"error"}</Typography>}
    //         <TableContainer component={Paper} sx={{ width: "50%" }}>
    //             <Table sx={{ minWidth: 650 }} aria-label="simple table" size="small">
    //                 <caption>Last observations</caption>
    //                 <TableHead>
    //                     <TableRow>
    //                         <TableCell>Value</TableCell>
    //                         <TableCell>Instant</TableCell>
    //                     </TableRow>
    //                 </TableHead>
    //                 <TableBody>
    //                     {observations.map((ob: observation) => (
    //                         <TableRow
    //                             key={ob.id}
    //                             sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
    //                         >
    //                             <TableCell>{`${ob.value} ${getObservationUnit(ob)}`}</TableCell>
    //                             <TableCell>{formatLastObservationInstant(ob)}</TableCell>
    //                         </TableRow>
    //                     ))}
    //                 </TableBody>
    //             </Table>
    //         </TableContainer>
    //     </>
    // )
    return (
        <></>
    )
}