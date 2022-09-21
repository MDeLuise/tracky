import { Dialog, DialogTitle, DialogContent, DialogActions, Button, TextField } from "@mui/material";
import { tracker } from "../interfaces";
import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";

export default function EditTracker(props: {
    open: boolean,
    tracker?: tracker,
    close: () => void,
    requestor: AxiosInstance,
    onSuccess: () => void,
    onFail: (error: string) => void
}) {
    const [trackerName, setTrackerName] = useState<string>("");
    const [trackerDescription, setTrackerDescription] = useState<string>("");
    const [trackerUnit, setTrackerUnit] = useState<string>("");

    const editTracker = (): void => {
        props.requestor.put(`/tracker/${props.tracker?.id}`, {
            name: trackerName,
            description: trackerDescription,
            unit: trackerUnit
        })
            .then((_res) => {
                props.onSuccess();
            })
            .catch((error) => {
                props.onFail(error);
            })
    };

    useEffect(() => {
        if (props.tracker != undefined) {
            setTrackerName(props.tracker.name);
            setTrackerDescription(props.tracker.description);
            setTrackerUnit(props.tracker.unit);
        }
    }, [props.tracker])


    return (
        <Dialog
            open={props.open}
            onClose={() => console.log("")}
            aria-labelledby="responsive-dialog-title"
        >
            <DialogTitle id="responsive-dialog-title">
                Edit a Tracker
            </DialogTitle>
            <DialogContent>
                <TextField
                    autoFocus
                    margin="dense"
                    label="Name"
                    type="text"
                    fullWidth
                    variant="standard"
                    defaultValue={trackerName}
                    onChange={(target) => setTrackerName(target.currentTarget.value)}
                />
                <TextField
                    autoFocus
                    margin="dense"
                    label="Description"
                    type="text"
                    fullWidth
                    variant="standard"
                    defaultValue={trackerDescription}
                    onChange={(target) => setTrackerDescription(target.currentTarget.value)}
                />
                <TextField
                    autoFocus
                    margin="dense"
                    label="Unit"
                    type="text"
                    fullWidth
                    variant="standard"
                    defaultValue={trackerUnit}
                    onChange={(target) => setTrackerUnit(target.currentTarget.value)}
                />
            </DialogContent>
            <DialogActions>
                <Button autoFocus onClick={props.close}>
                    Cancel
                </Button>
                <Button onClick={() => {
                    props.close();
                    editTracker();
                }}
                    autoFocus>
                    Save
                </Button>
            </DialogActions>
        </Dialog>
    )
}