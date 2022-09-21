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
            ...props.style,
            ...{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                gap: "30px",
                marginBottom: "20px",
                flexWrap: "wrap"
            }
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