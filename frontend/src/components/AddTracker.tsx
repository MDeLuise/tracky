import { AxiosInstance } from "axios";
import { Box, Button, TextField } from "@mui/material";
import AddIcon from '@mui/icons-material/Add';
import { useState } from "react";

export default function AddTracker(props: {
    requestor: AxiosInstance,
    onSuccess?: () => void,
    style?: {}
}) {
    const [trackerName, setTrackerName] = useState<string>("");
    const [trackerDescription, setTrackerDescription] = useState<string>("");
    const [trackeUnit, setTrackerUnit] = useState<string>("");
    const [error, setError] = useState<string>();

    const addTracker = (): void => {
        props.requestor.post("/tracker", {
            name: trackerName,
            description: trackerDescription,
            unit: trackeUnit
        })
            .then((_response) => {
                if (props.onSuccess != undefined) {
                    props.onSuccess()
                }
            })
            .catch((error) => setError(error))
    };

    return (
        <Box sx={{
            maxWidth: "fit-content",
            display: "flex",
            backgroundImage: "linear-gradient(rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0.05))",
            padding: "15px",
            borderRadius: "10px",
            alignItems: "center",
            gap: "30px",
            margin: "30px auto",
            flexWrap: "wrap",
            ...props.style,
        }}>
            <TextField label="Name" variant="standard" type="text" onChange={(target) => setTrackerName(target.currentTarget.value)} />
            <TextField label="Description" variant="standard" type="text" onChange={(target) => setTrackerDescription(target.currentTarget.value)} />
            <TextField label="Unit" variant="standard" type="text" onChange={(target) => setTrackerUnit(target.currentTarget.value)} />
            <Button variant="contained" endIcon={<AddIcon />} onClick={addTracker}>
                Add
            </Button>
        </Box>
    )
}