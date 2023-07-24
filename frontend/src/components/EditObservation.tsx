import { Dialog, DialogTitle, DialogContent, DialogActions, Button, TextField, NativeSelect, FormControl, InputLabel } from "@mui/material";
import { observation, tracker } from "../interfaces";
import { AxiosInstance } from "axios";
import { DateTimePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { useEffect, useState } from "react";
import dayjs, { Dayjs } from "dayjs";

export default function EditObservation(props: {
    open: boolean,
    trackers: tracker[],
    close: () => void,
    observation?: observation,
    requestor: AxiosInstance,
    onSuccess: () => void,
    onFail: (error: string) => void;
}) {
    const [observationValue, setObservationValue] = useState<number>(0);
    const [observationNote, setObservationNote] = useState<string>();
    const [observationDate, setObservationDate] = useState<Dayjs>(dayjs(new Date()));
    const [observationTrackerId, setObservationTrackerId] = useState<number>(0);

    const editObservation = (): void => {
        props.requestor.put(`/observation/${props.observation?.id}`, {
            value: observationValue,
            trackerId: observationTrackerId,
            instant: observationDate,
            note: observationNote,
        })
            .then((_res) => {
                props.onSuccess();
            })
            .catch((error) => {
                props.onFail(error);
            });
    };

    useEffect(() => {
        if (props.observation != undefined) {
            setObservationValue(props.observation.value);
            setObservationDate(dayjs(props.observation.instant));
            setObservationTrackerId(props.observation.trackerId);
            setObservationNote(props.observation.note);
        }
    }, [props.observation]);


    return (
        <Dialog
            open={props.open}
            aria-labelledby="responsive-dialog-title"
        >
            <DialogTitle id="responsive-dialog-title">
                Edit an observation
            </DialogTitle>
            <DialogContent>
                <FormControl fullWidth>
                    <InputLabel variant="standard" htmlFor="uncontrolled-native">
                        Tracker
                    </InputLabel>
                    <NativeSelect
                        defaultValue={observationTrackerId}
                        inputProps={{
                            name: 'tracker'
                        }}
                        onChange={(target) => setObservationTrackerId(Number(target.currentTarget.value))}
                    >
                        {props.trackers.map((tr: tracker) => {
                            return <option key={tr.id} value={tr.id}>{tr.name}</option>;
                        })}
                    </NativeSelect>
                </FormControl>
                <TextField
                    autoFocus
                    margin="dense"
                    label="Value"
                    type="number"
                    fullWidth
                    variant="standard"
                    defaultValue={observationValue}
                    onChange={(target) => setObservationValue(Number(target.currentTarget.value))}
                />
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DateTimePicker
                        defaultValue={dayjs(observationDate)}
                        onChange={(newValue) => setObservationDate(newValue != null ? newValue : dayjs(new Date()))}
                        format="DD/MM/YYYY, HH:mm"
                    />
                </LocalizationProvider>
                <TextField
                    autoFocus
                    margin="dense"
                    label="Note"
                    type="text"
                    fullWidth
                    variant="standard"
                    defaultValue={observationNote}
                    onChange={(target) => setObservationNote(target.currentTarget.value)}
                />
            </DialogContent>
            <DialogActions>
                <Button autoFocus onClick={props.close}>
                    Cancel
                </Button>
                <Button onClick={() => {
                    props.close();
                    editObservation();
                }}
                    autoFocus>
                    Save
                </Button>
            </DialogActions>
        </Dialog>
    );
}