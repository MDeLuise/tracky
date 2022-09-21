import { AxiosInstance } from "axios";
import { tracker } from "../interfaces";
import { Box, Button, NativeSelect, TextField } from "@mui/material";
import { LocalizationProvider, DateTimePicker } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import AddIcon from '@mui/icons-material/Add';
import dayjs, { Dayjs } from "dayjs";
import { useEffect, useState } from "react";

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
            ...props.style,
            ...{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                gap: "30px",
                marginBottom: "20px",
                flexWrap: "wrap",
                // border: 0.5,
                // borderRadius: 1,
                // borderColor: "primary.secondary"
            }
        }}>
            {
                props.trackerId == undefined &&
                (<NativeSelect
                    defaultValue={props.trackerId}
                    inputProps={{
                        name: 'tracker'
                    }}
                    onChange={(target) => setTrackerId(Number(target.currentTarget.value))}
                    sx={{ marginTop: "15px" }}
                >
                    {trackers.map((tr: tracker) => {
                        return <option
                            value={tr.id}
                            key={tr.id}>
                            {tr.name}
                        </option>;
                    })}
                </NativeSelect>
                )
            }
            <TextField label="Value" variant="standard" type="number" onChange={(target) => setValue(Number(target.currentTarget.value))} />
            <LocalizationProvider dateAdapter={AdapterDayjs}>
                <DateTimePicker defaultValue={dayjs(new Date())} onChange={(newValue) => setDate(newValue != null ? newValue : dayjs(new Date()))} />
            </LocalizationProvider>
            <TextField label="Note" variant="standard" type="text" onChange={(target) => setNote(target.currentTarget.value)} />
            <Button variant="contained" endIcon={<AddIcon />} onClick={addObservation}>
                Add
            </Button>
        </Box>
    );
}