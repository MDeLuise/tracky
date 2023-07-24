import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";
import { observation, tracker } from "../interfaces";
import { useParams } from "react-router-dom";
import { Alert, Box, Button, Checkbox, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton, Paper, Snackbar, Table, TableBody, TableCell, TableContainer, TableFooter, TableHead, TablePagination, TableRow, Toolbar, Tooltip, Typography } from "@mui/material";
import Navbar from "./Navbar";
import Chart, { CategoryScale } from "chart.js/auto";
import LineChart from "./LineChart";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import { alpha } from '@mui/material/styles';
import EditObservation from "./EditObservation";
import AddObservation from "./AddObservation";
import { isSmallScreen } from "../common";

function ConfirmationBox(props: {
    open: boolean,
    selected: number,
    close: () => void,
    clearSelected: () => void,
    removeSelected: () => void;
}) {
    return (
        <Dialog
            open={props.open}
            onClose={props.close}
            aria-labelledby="responsive-dialog-title"
        >
            <DialogTitle id="responsive-dialog-title">
                {"Confirm delete"}
            </DialogTitle>
            <DialogContent>
                <DialogContentText>
                    Are you sure you want to remove the selected observations?
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button autoFocus onClick={props.close}>
                    No
                </Button>
                <Button onClick={() => {
                    props.close();
                    props.removeSelected();
                    props.clearSelected();
                }}
                    autoFocus>
                    Yes
                </Button>
            </DialogActions>
        </Dialog>
    );
};

function ObservationTableToolbar(props: {
    selected: number,
    openDeleteConfirmDialog: () => void,
    openEditDialog: () => void,
    onEditSuccess: () => void,
    onEditFails: (err: string) => void;
}) {
    return (
        <Toolbar
            sx={{
                pl: { sm: 2 },
                pr: { xs: 1, sm: 1 },
                ...(props.selected > 0 && {
                    bgcolor: (theme) =>
                        alpha(theme.palette.primary.main, theme.palette.action.activatedOpacity),
                }),
            }}
        >
            {props.selected > 0 ? (
                <Typography
                    sx={{ flex: '1 1 100%' }}
                    color="inherit"
                    variant="subtitle1"
                    component="div"
                >
                    {props.selected} selected
                </Typography>
            ) : (
                <Typography
                    sx={{ flex: '1 1 100%' }}
                    variant="h6"
                    id="tableTitle"
                    component="div"
                >
                    Observations
                </Typography>
            )}
            {props.selected == 1 &&
                <Tooltip title="Edit">
                    <IconButton>
                        <EditIcon onClick={props.openEditDialog} />
                    </IconButton>
                </Tooltip>
            }
            {props.selected > 0 &&
                <Tooltip title="Delete">
                    <IconButton>
                        <DeleteIcon onClick={props.openDeleteConfirmDialog} />
                    </IconButton>
                </Tooltip>
            }
        </Toolbar>
    );
}

function ObservationTableBody(props: {
    totalElements: number,
    requireElementsPerPage: number,
    pageNo: number,
    fetchObservations: (getPageNo: number, getElementsPerPage: number) => void,
    unit: string,
    observations: observation[],
    selectedIds: Set<number>,
    setSelectedIds: (selected: Set<number>) => void,
    setObservationToEdit: (ob: observation) => void;
}) {
    const formatLastObservationInstant = (ob: observation): string => {
        let lastObservationDate = new Date(ob.instant).toLocaleDateString();
        let lastObservationTime = new Date(ob.instant).toLocaleTimeString();
        return `${lastObservationDate}, ${lastObservationTime}`;
    };
    return (
        <Table aria-label="List of tracker's observations" size="small">
            <TableFooter>
                <TableRow>
                    <TablePagination
                        rowsPerPageOptions={[25, 50, 100, { label: 'All', value: -1 }]}
                        colSpan={5}
                        count={props.totalElements}
                        rowsPerPage={props.requireElementsPerPage == props.totalElements ? -1 : props.requireElementsPerPage}
                        page={props.pageNo}
                        SelectProps={{
                            inputProps: {
                                'aria-label': 'rows per page',
                            },
                            native: true,
                        }}
                        onPageChange={(_event: React.MouseEvent<HTMLButtonElement> | null, newPage: number,) => {
                            props.fetchObservations(newPage, props.requireElementsPerPage);
                        }}
                        onRowsPerPageChange={(event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                            let newElementsNumber: number = parseInt(event.target.value, 10);
                            props.fetchObservations(0, newElementsNumber != -1 ? newElementsNumber : props.totalElements);
                        }}
                    />
                </TableRow>
            </TableFooter>
            <TableHead>
                <TableRow>
                    <TableCell align="center">
                        <Checkbox
                            color="primary"
                            inputProps={{
                                'aria-label': 'select all observations',
                            }}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                if (event.target.checked) {
                                    props.setSelectedIds(new Set([...props.selectedIds, ...props.observations.map(ob => ob.id)]));
                                } else {
                                    props.setSelectedIds(new Set());
                                }
                            }}
                            checked={props.observations.length > 0 && props.observations.length == props.selectedIds.size}
                            indeterminate={props.selectedIds.size > 0 && props.selectedIds.size < Math.min(props.totalElements, props.requireElementsPerPage)}
                            disabled={props.totalElements == 0}
                        />
                    </TableCell>
                    <TableCell align="center">Value</TableCell>
                    <TableCell align="center">Date</TableCell>
                    <TableCell align="center">Note</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {props.observations.map((ob: observation) => (
                    <TableRow
                        key={ob.id}
                        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                    >
                        <TableCell padding="checkbox" align="center">
                            <Checkbox
                                color="primary"
                                inputProps={{
                                    'aria-label': 'select observation',
                                }}
                                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                    if (event.target.checked) {
                                        props.setSelectedIds(new Set([...props.selectedIds, ob.id]));
                                        props.setObservationToEdit(ob);
                                    } else {
                                        props.setSelectedIds(new Set([...props.selectedIds].filter((id: number) => id !== ob.id)));
                                    }
                                }}
                                checked={props.selectedIds.has(ob.id)}
                            />
                        </TableCell>
                        <TableCell align="center">{`${ob.value} ${props.unit}`}</TableCell>
                        <TableCell align="center">{formatLastObservationInstant(ob)}</TableCell>
                        <TableCell align="center">{ob.note}</TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>);
}

function ObservationTable(props: {
    totalElements: number,
    requireElementsPerPage: number,
    pageNo: number,
    fetchObservations: (getPageNo: number, getElementsPerPage: number) => void,
    unit: string,
    observations: observation[],
    selectedIds: Set<number>,
    setSelectedIds: (selected: Set<number>) => void,
    openDeleteConfirmDialog: () => void,
    requestor: AxiosInstance,
    trackers: tracker[],
    onEditSuccess: () => void,
    onEditFails: (err: string) => void;
}) {
    const [openEditDialog, setOpenEditDialog] = useState<boolean>(false);
    const [observationToEdit, setObservationToEdit] = useState<observation | undefined>();
    return (
        <>
            <EditObservation
                open={openEditDialog}
                trackers={props.trackers}
                close={() => setOpenEditDialog(false)}
                requestor={props.requestor}
                observation={observationToEdit}
                onSuccess={props.onEditSuccess}
                onFail={props.onEditFails}
            />
            <Box sx={{ overflow: "auto" }}>
                <Box sx={{ width: "90%", display: "table", tableLayout: "fixed", margin: "0 auto" }}>
                    <TableContainer component={Paper} sx={{ minWidth: "80vw", margin: "0 auto" }}>
                        <ObservationTableToolbar
                            selected={props.selectedIds.size}
                            openDeleteConfirmDialog={props.openDeleteConfirmDialog}
                            openEditDialog={() => setOpenEditDialog(true)}
                            onEditSuccess={props.onEditSuccess}
                            onEditFails={props.onEditFails}
                        />
                        <ObservationTableBody
                            fetchObservations={props.fetchObservations}
                            observations={props.observations}
                            pageNo={props.pageNo}
                            requireElementsPerPage={props.requireElementsPerPage}
                            totalElements={props.totalElements}
                            unit={props.unit}
                            setSelectedIds={props.setSelectedIds}
                            selectedIds={props.selectedIds}
                            setObservationToEdit={setObservationToEdit}
                        />
                    </TableContainer>
                </Box>
            </Box>
        </>
    );
}

export default function Tracker(props: {
    isLoggedIn: () => boolean,
    requestor: AxiosInstance;
}) {
    const { trackerId } = useParams();
    const [error, setError] = useState<string>();
    const [observations, setObservations] = useState<observation[]>([]);
    const [pageNo, setPageNo] = useState<number>(0);
    const [requireElementsPerPage, setRequireElementsPerPage] = useState<number>(25);
    const [totalElements, setTotalElements] = useState<number>(0);
    const [selectedObservationIds, setSelectedObservationIds] = useState<Set<number>>(new Set());
    const [openDeleteConfirmationDialog, setOpenDeleteConfirmationDialog] = useState<boolean>(false);
    const [alertSeverity, setAlertSeverity] = useState<"error" | "warning" | "info" | "success">("success");
    const [alertMessage, setAlertMessage] = useState<string>();
    const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
    const [trackers, setTrackers] = useState<tracker[]>([]);

    Chart.register(CategoryScale);

    const getObservation = (getPageNo: number, getElementsPerPage: number): void => {
        setPageNo(getPageNo);
        setRequireElementsPerPage(getElementsPerPage);
        props.requestor.get(`observation?trackerId=${trackerId}&pageNo=${getPageNo}&pageSize=${getElementsPerPage}&sortBy=instant&sortDir=DESC`)
            .then(response => {
                let fetchedObservations: observation[] = [];
                response.data.content.forEach((ob: observation) => {
                    fetchedObservations.push(ob);
                });
                setObservations(fetchedObservations);
                setTotalElements(response.data.totalElements);
            })
            .catch(error => setError(error));
    };

    const removeSelectedObservations = (): void => {
        let joinedIds: string = Array.from(selectedObservationIds).join(",");
        props.requestor.delete(`/observation/${joinedIds}`)
            .then(_response => {
                setAlertMessage("Observation(s) deleted successfully");
                setAlertSeverity("success");
                setSnackbarOpen(true);
                setObservations(observations.filter((ob: observation) => !selectedObservationIds.has(ob.id)));
            })
            .catch(error => {
                setError(error);
                setAlertMessage("Error while deleting observation(s)");
                setAlertSeverity("error");
            });
    };

    const editSuccess = (): void => {
        setAlertMessage("Observation updated successfully");
        setAlertSeverity("success");
        setSnackbarOpen(true);
        getObservation(pageNo, requireElementsPerPage);
    };

    const editFail = (error: string): void => {
        setAlertMessage("Error while updating observation");
        setAlertSeverity("error");
        setSnackbarOpen(true);
        console.error(error);
    };

    const handleClose = (_event?: React.SyntheticEvent | Event, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setSnackbarOpen(false);
    };

    const getTrackers = (): void => {
        props.requestor.get("tracker/count")
            .then(response => {
                getAllTrackers(response.data);
            })
            .catch(error => setError(error));
    };

    const getAllTrackers = (total: number): void => {
        props.requestor.get(`tracker?pageSize=${total}`)
            .then(response => {
                let fetchedTrackers: tracker[] = [];
                response.data.content.forEach((ft: tracker) => {
                    fetchedTrackers.push(ft);
                });
                setTrackers(fetchedTrackers);
            })
            .catch(error => setError(error));
    };

    useEffect(() => {
        getObservation(0, 25);
        getTrackers();
    }, []);

    return (
        <>
            <Navbar requestor={props.requestor} mobile={isSmallScreen()} />
            {error && <Typography variant="body1">{"error"}</Typography> ||
                <>
                    <LineChart observations={observations.slice().reverse()} />

                    <Snackbar open={snackbarOpen} autoHideDuration={3000} onClose={handleClose}>
                        <Alert variant="filled" severity={alertSeverity} >
                            {alertMessage}
                        </Alert>
                    </Snackbar>

                    <ConfirmationBox
                        selected={selectedObservationIds.size}
                        open={openDeleteConfirmationDialog}
                        close={() => setOpenDeleteConfirmationDialog(false)}
                        clearSelected={() => setSelectedObservationIds(new Set())}
                        removeSelected={removeSelectedObservations}
                    />

                    <AddObservation
                        requestor={props.requestor}
                        trackerId={Number(trackerId)}
                        onSuccess={() => getObservation(0, 25)}
                        style={{
                            marginTop: "30px",
                            width: "90% !important",
                        }}
                    />
                    <ObservationTable
                        fetchObservations={getObservation}
                        observations={observations}
                        pageNo={pageNo}
                        requireElementsPerPage={requireElementsPerPage}
                        totalElements={totalElements}
                        unit={observations.length > 0 ? observations[0].unit : ""}
                        setSelectedIds={setSelectedObservationIds}
                        selectedIds={selectedObservationIds}
                        openDeleteConfirmDialog={() => setOpenDeleteConfirmationDialog(true)}
                        requestor={props.requestor}
                        trackers={trackers}
                        onEditSuccess={editSuccess}
                        onEditFails={editFail}
                    />
                </>
            }
        </>
    );
}