import { Typography, TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody } from "@mui/material";
import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";

export default function AccountStats(props: { requestor: AxiosInstance, newData: number }) {
    const [stats, setStats] = useState<{}>({});
    const [error, setError] = useState<string>();

    const getStats = (): void => {
        props.requestor.get("account-stats")
            .then((res) => setStats(res.data))
            .catch((error) => setError(error))
    }

    useEffect(() => {
        getStats();
    }, [props.newData]);

    return (
        <>
            {error && <Typography variant="body1">{"error"}</Typography>}
            <TableContainer component={Paper} sx={{ width: "100%" }}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table" size="small">
                    <caption>Account statistics</caption>
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell>Value</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {Object.entries(stats).map((keyValArray) => (
                            <TableRow
                                key={keyValArray[0]}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell>{keyValArray[0]}</TableCell>
                                <TableCell>{keyValArray[1] as string}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}