import { Box, BottomNavigation, BottomNavigationAction, Paper } from "@mui/material";
import { useState } from "react";
import InsertChartIcon from '@mui/icons-material/InsertChart';
import LabelIcon from '@mui/icons-material/Label';
import { AxiosInstance } from "axios";
import Navbar from "./Navbar";
import { isSmallScreen } from "../common";

export default function AddEntities(props: { requestor: AxiosInstance, isLoggedIn: () => boolean }) {
    const [value, setValue] = useState(0);

    return (
        <>
            <Navbar requestor={props.requestor} mobile={isSmallScreen()} />
            <Paper sx={{ position: 'fixed', bottom: 0, left: 0, right: 0 }} elevation={3}>
                <BottomNavigation
                    showLabels
                    value={value}
                    onChange={(event, newValue) => {
                        setValue(newValue);
                    }}
                >
                    <BottomNavigationAction label="Tracker" icon={<InsertChartIcon />} />
                    <BottomNavigationAction label="Observation" icon={<LabelIcon />} />
                </BottomNavigation>
            </Paper>
        </>
    )
}