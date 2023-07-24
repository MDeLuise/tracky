import { AxiosInstance } from "axios";
import { useEffect, useState } from "react";
import { tracker } from "../interfaces";
import Navbar from "./Navbar";
import { Typography, TableContainer, Paper, Table, TableHead, TableRow, TableCell, TableBody, Alert, Box, Button, Checkbox, Dialog, DialogActions, DialogContent, DialogContentText, DialogTitle, IconButton, Snackbar, TableFooter, TablePagination, Toolbar, Tooltip, alpha } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import EditTracker from "./EditTracker";
import AddTracker from "./AddTracker";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { isSmallScreen } from "../common";

function ConfirmationBox(props: {
    open: boolean,
    selected: number,
    close: () => void,
    clearSelected: () => void,
    removeSelected: () => void
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
                    Are you sure you want to remove the selected trackers?
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
    )
};

function TrackerTableToolbar(props: {
    selected: number,
    openDeleteConfirmDialog: () => void,
    openEditDialog: () => void
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
                    Trackers
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
    )
}

function TrackerTableBody(props: {
    totalElements: number,
    requireElementsPerPage: number,
    pageNo: number,
    fetchTrackers: (getPageNo: number, getElementsPerPage: number) => void,
    unit: string,
    trackers: tracker[],
    selectedIds: Set<number>,
    setSelectedIds: (selected: Set<number>) => void,
    setTrackerToEdit: (tr: tracker) => void
}) {
    let navigate: NavigateFunction = useNavigate();
    const formatLastUpdatedInstant = (tr: tracker | undefined): string => {
        if (tr == undefined) {
            return "";
        }
        let lastObservationInstant: string = tr.lastObservation != undefined ?
            tr.lastObservation.instant.toString() :
            "";
        let lastObservationDate = new Date(lastObservationInstant).toLocaleDateString();
        let lastObservationTime = new Date(lastObservationInstant).toLocaleTimeString();
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
                            native: false,
                        }}
                        onPageChange={(_event: React.MouseEvent<HTMLButtonElement> | null, newPage: number,) => {
                            props.fetchTrackers(newPage, props.requireElementsPerPage);
                        }}
                        onRowsPerPageChange={(event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
                            let newElementsNumber: number = parseInt(event.target.value, 10);
                            props.fetchTrackers(0, newElementsNumber != -1 ? newElementsNumber : props.totalElements);
                        }}
                    />
                </TableRow>
            </TableFooter>
            <TableHead>
                <TableRow>
                    <TableCell align="center">
                        <Checkbox
                            color="primary"
                            indeterminate={props.selectedIds.size > 0 && props.selectedIds.size < Math.min(props.totalElements, props.requireElementsPerPage)}
                            disabled={props.totalElements == 0}
                            inputProps={{
                                'aria-label': 'select all observations',
                            }}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                if (event.target.checked) {
                                    props.setSelectedIds(new Set([...props.selectedIds, ...props.trackers.map(ob => ob.id)]));
                                } else {
                                    props.setSelectedIds(new Set());
                                }
                            }}
                            checked={props.trackers.length == props.selectedIds.size}
                        />
                    </TableCell>
                    <TableCell align="center">Name</TableCell>
                    <TableCell align="center">Description</TableCell>
                    <TableCell align="center">Last value</TableCell>
                    <TableCell align="center">Updated on</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {props.trackers.map((tr: tracker) => (
                    <TableRow
                        key={tr.id}
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
                                        props.setSelectedIds(new Set([...props.selectedIds, tr.id]));
                                        props.setTrackerToEdit(tr);
                                    } else {
                                        props.setSelectedIds(new Set([...props.selectedIds].filter((id: number) => id !== tr.id)));
                                    }
                                }}
                                checked={props.selectedIds.has(tr.id)}
                            />
                        </TableCell>
                        <TableCell
                            align="center"
                            onClick={() => navigate(`/tracker/${tr.id}`)}
                            className="clickable">{tr.name}
                        </TableCell>
                        <TableCell
                            align="center"
                            onClick={() => navigate(`/tracker/${tr.id}`)}
                            className="clickable">{tr.description}
                        </TableCell>
                        <TableCell
                            align="center"
                            onClick={() => navigate(`/tracker/${tr.id}`)}
                            className="clickable">
                            {tr.lastObservation != undefined ? tr.lastObservation.value + " " + tr.unit : ""}
                        </TableCell>
                        <TableCell
                            align="center"
                            onClick={() => navigate(`/tracker/${tr.id}`)}
                            className="clickable">
                            {formatLastUpdatedInstant(tr.lastObservation != undefined ? tr : undefined)}
                        </TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>)
}

function TrackerTable(props: {
    totalElements: number,
    requireElementsPerPage: number,
    pageNo: number,
    fetchTrackers: (getPageNo: number, getElementsPerPage: number) => void,
    unit: string,
    trackers: tracker[],
    selectedIds: Set<number>,
    setSelectedIds: (selected: Set<number>) => void,
    openDeleteConfirmDialog: () => void,
    requestor: AxiosInstance,
    onEditSuccess: () => void,
    onEditFails: (err: string) => void
}) {
    const [openEditDialog, setOpenEditDialog] = useState<boolean>(false);
    const [trackerToEdit, setTrackerToEdit] = useState<tracker | undefined>();
    return (
        <>
            <EditTracker
                open={openEditDialog}
                close={() => setOpenEditDialog(false)}
                requestor={props.requestor}
                tracker={trackerToEdit}
                onSuccess={props.onEditSuccess}
                onFail={props.onEditFails}
            />
            <Box sx={{ overflow: "auto" }}>
                <Box sx={{ width: "90%", display: "table", tableLayout: "fixed", margin: "0 auto" }}>
                    <TableContainer component={Paper} sx={{ minWidth: "80vw", margin: "0 auto" }}>
                        <TrackerTableToolbar
                            selected={props.selectedIds.size}
                            openDeleteConfirmDialog={props.openDeleteConfirmDialog}
                            openEditDialog={() => setOpenEditDialog(true)}
                        />
                        <TrackerTableBody
                            fetchTrackers={props.fetchTrackers}
                            trackers={props.trackers}
                            pageNo={props.pageNo}
                            requireElementsPerPage={props.requireElementsPerPage}
                            totalElements={props.totalElements}
                            unit={props.unit}
                            setSelectedIds={props.setSelectedIds}
                            selectedIds={props.selectedIds}
                            setTrackerToEdit={setTrackerToEdit}
                        />
                    </TableContainer>
                </Box>
            </Box>
        </>
    )
}

export default function AllTrackers(props: {
    isLoggedIn: () => boolean,
    requestor: AxiosInstance
}) {
    const [error, setError] = useState<string>();
    const [unit, setUnit] = useState<string>("");
    const [pageNo, setPageNo] = useState<number>(0);
    const [requireElementsPerPage, setRequireElementsPerPage] = useState<number>(25);
    const [totalElements, setTotalElements] = useState<number>(0);
    const [selectedTrackerIds, setSelectedTrackerIds] = useState<Set<number>>(new Set());
    const [openDeleteConfirmationDialog, setOpenDeleteConfirmationDialog] = useState<boolean>(false);
    const [alertSeverity, setAlertSeverity] = useState<"error" | "warning" | "info" | "success">("success");
    const [alertMessage, setAlertMessage] = useState<string>();
    const [snackbarOpen, setSnackbarOpen] = useState<boolean>(false);
    const [trackers, setTrackers] = useState<tracker[]>([]);

    const removeSelectedTrackers = (): void => {
        let joinedIds: string = Array.from(selectedTrackerIds).join(",");
        props.requestor.delete(`/tracker/${joinedIds}`)
            .then(_response => {
                setAlertMessage("Tracker(s) deleted successfully");
                setAlertSeverity("success");
                setSnackbarOpen(true);
                setTrackers(trackers.filter((tr: tracker) => !selectedTrackerIds.has(tr.id)));
            })
            .catch(error => {
                setError(error);
                setAlertMessage("Error while deleting tracker(s)");
                setAlertSeverity("error");
            })
    };

    const handleClose = (_event?: React.SyntheticEvent | Event, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setSnackbarOpen(false);
    };

    const getTrackers = (getPageNo: number, getElementsPerPage: number): void => {
        setPageNo(getPageNo);
        setRequireElementsPerPage(getElementsPerPage);
        props.requestor.get(`tracker?pageNo=${getPageNo}&pageSize=${getElementsPerPage}`)
            .then(response => {
                let fetchedTrackers: tracker[] = [];
                response.data.content.forEach((ft: tracker) => {
                    fetchedTrackers.push(ft);
                });
                setTrackers(fetchedTrackers);
                setTotalElements(response.data.totalElements);
            })
            .catch(error => setError(error))
    };

    const editSuccess = (): void => {
        setAlertMessage("Target updated successfully");
        setAlertSeverity("success");
        setSnackbarOpen(true);
        getTrackers(pageNo, requireElementsPerPage);
    };

    const editFail = (error: string): void => {
        setAlertMessage("Error while updating target");
        setAlertSeverity("error");
        setSnackbarOpen(true);
        console.error(error);
    };

    useEffect(() => {
        getTrackers(pageNo, requireElementsPerPage);
    }, []);

    return (
        <>
            <Navbar requestor={props.requestor} mobile={isSmallScreen()} />
            {error && <Typography variant="body1">{"error"}</Typography> ||
                <>
                    <Snackbar open={snackbarOpen} autoHideDuration={3000} onClose={handleClose}>
                        <Alert variant="filled" severity={alertSeverity} >
                            {alertMessage}
                        </Alert>
                    </Snackbar>

                    <ConfirmationBox
                        selected={selectedTrackerIds.size}
                        open={openDeleteConfirmationDialog}
                        close={() => setOpenDeleteConfirmationDialog(false)}
                        clearSelected={() => setSelectedTrackerIds(new Set())}
                        removeSelected={removeSelectedTrackers}
                    />

                    <AddTracker
                        requestor={props.requestor}
                        onSuccess={() => getTrackers(pageNo, requireElementsPerPage)}
                        style={{ marginTop: "30px" }}
                    />
                    <TrackerTable
                        fetchTrackers={getTrackers}
                        pageNo={pageNo}
                        requireElementsPerPage={requireElementsPerPage}
                        totalElements={totalElements}
                        unit={unit}
                        setSelectedIds={setSelectedTrackerIds}
                        selectedIds={selectedTrackerIds}
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