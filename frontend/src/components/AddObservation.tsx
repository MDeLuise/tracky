import { AxiosInstance } from "axios";
import { tracker } from "../interfaces";
import { Box, Button, MenuItem, NativeSelect, Select, TextField } from "@mui/material";
import { LocalizationProvider, DateTimePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import AddIcon from '@mui/icons-material/Add';
import dayjs, { Dayjs } from "dayjs";
import { useEffect, useState } from "react";
import { isSmallScreen } from "../common";

export default function AddObservation(props: {
    trackerId?: number,
    trackers?: tracker[],
    requestor: AxiosInstance,
    onSuccess?: () => void,
    style?: {};
}) {
    const [trackers, setTrackers] = useState<tracker[]>([]);
    const [trackerId, setTrackerId] = useState<number>(props.trackerId != undefined ? props.trackerId : -1);
    const [value, setValue] = useState<number>(0);
    const [date, setDate] = useState<Dayjs>(dayjs(new Date()));
    const [note, setNote] = useState<string>();
    const [error, setError] = useState<string>();

    const addObservation = (): void => {
        props.requestor.post("/observation", {
            trackerId: trackerId,
            value: value,
            instant: date,
            note: note,
        })
            .then((_response) => {
                if (props.onSuccess != undefined) {
                    props.onSuccess();
                }
            })
            .catch((error) => setError(error));
    };

    useEffect(() => {
        setTrackers(props.trackers != undefined ? props.trackers : []);
        if (trackerId === -1 && props.trackers != undefined && props.trackers.length > 0) {
            setTrackerId(props.trackers[0].id);
        }
    }, [props.trackers]);

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
            flexWrap: isSmallScreen() ? "wrap" : "nowrap",
            ...props.style,
        }}>
            {
                props.trackerId == undefined &&
                <Select
                    value={trackerId}
                    onChange={(event) => setTrackerId(Number(event.target.value))}
                    sx={{
                        marginTop: "15px",
                        width: isSmallScreen() ? "45%" : "15%",
                    }}
                >
                    {
                        trackers.map((tr: tracker) => {
                            return <MenuItem
                                value={tr.id}
                                key={tr.id}>
                                {tr.name}
                            </MenuItem>;
                        })
                    }
                </Select>
            }
            <TextField
                label="Value"
                variant="standard"
                type="number"
                onChange={(target) => setValue(Number(target.currentTarget.value))}
                sx={{
                    width: isSmallScreen() ? "45%" : "initial",
                }}
            />
            <LocalizationProvider
                dateAdapter={AdapterDayjs}
            >
                <DateTimePicker
                    format="DD/MM/YYYY, HH:mm"
                    defaultValue={dayjs(new Date())}
                    onChange={(newValue) => setDate(newValue != null ? newValue : dayjs(new Date()))}
                    sx={{
                        width: isSmallScreen() ? (props.trackerId != undefined ? "45%" : "100%") : "initial",
                    }}
                />
            </LocalizationProvider>
            <TextField
                label="Note"
                variant="standard"
                type="text"
                onChange={(target) => setNote(target.currentTarget.value)}
                sx={{
                    width: isSmallScreen() ? "100%" : "initial",
                }} />
            <Button
                variant="contained"
                endIcon={<AddIcon />}
                onClick={addObservation}
                sx={{
                    width: isSmallScreen() ? "100%" : "initial",
                }}>
                Add
            </Button>
        </Box>
    );
}