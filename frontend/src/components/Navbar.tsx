import * as React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import Button from '@mui/material/Button';
import TimelineOutlinedIcon from '@mui/icons-material/TimelineOutlined';
import MenuItem from '@mui/material/MenuItem';
import { AxiosInstance } from 'axios';
import { NavigateFunction, useNavigate } from 'react-router-dom';
import secureLocalStorage from 'react-secure-storage';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import { Link } from '@mui/material';
import SettingsOutlinedIcon from '@mui/icons-material/SettingsOutlined';


export default function Navbar(props: {
    mobile: boolean,
    requestor: AxiosInstance,
}) {
    const [anchorElNav, setAnchorElNav] = React.useState<null | HTMLElement>(null);
    const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);
    const username: string = secureLocalStorage.getItem("tracky-username") as string;
    let navigate: NavigateFunction = useNavigate();

    const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>): void => {
        setAnchorElNav(event.currentTarget);
    };

    const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>): void => {
        setAnchorElUser(event.currentTarget);
    };

    const handleCloseNavMenu = (): void => {
        setAnchorElNav(null);
    };

    const handleCloseUserMenu = (): void => {
        setAnchorElUser(null);
    };

    const logout = (): void => {
        handleCloseUserMenu();
        secureLocalStorage.removeItem("tracky-key");
        secureLocalStorage.removeItem("tracky-username");
        navigate("/auth");
    };

    return (
        <Box
            sx={{
                display: "flex",
                flexDirection: "row",
                width: "100vw",
                padding: props.mobile ? "10px 15px" : "15px 50px",
                gap: props.mobile? "10px" : "20px",
                alignItems: "center",
                backgroundColor: "secondary.main",
                '& a': {
                    textDecoration: "none",
                    color: "inherit"
                },
            }}>
            <Link
                href='/'
                sx={{
                    display: "flex",
                    alignItems: "center",
                }}
            >
                <TimelineOutlinedIcon sx={{ marginRight: "5px" }} />
                <Typography
                    variant="h6"
                    noWrap
                    sx={{
                        fontWeight: 700,
                        letterSpacing: props.mobile ? '.1rem' : '.3rem',
                        color: 'inherit',
                        fontFamily: "Bebas Neue, sans-serif",
                    }}
                >
                    Tracky
                </Typography>
            </Link>
            <Link href='/trackers'>
                {props.mobile ? "trackers" : "All trackers"}
            </Link>
            <Link href='/observations'>
                {props.mobile ? "observations" : "All observations"}
            </Link>


            <Box
                sx={{
                    flexGrow: 1,
                    display: "flex",
                    justifyContent: "flex-end"
                }}
            >
                <Button
                    endIcon={<KeyboardArrowDownIcon />}
                    onClick={handleOpenUserMenu}
                    sx={{
                        textTransform: "none",
                        color: "inherit",
                    }}
                >
                    {props.mobile ? < SettingsOutlinedIcon /> : username}
                </Button>
                <Menu
                    sx={{ mt: '45px' }}
                    id="menu-appbar"
                    anchorEl={anchorElUser}
                    anchorOrigin={{
                        vertical: 'top',
                        horizontal: 'right',
                    }}
                    keepMounted
                    transformOrigin={{
                        vertical: 'top',
                        horizontal: 'right',
                    }}
                    open={Boolean(anchorElUser)}
                    onClose={handleCloseUserMenu}
                >
                    {/* <MenuItem key="settings" onClick={logout}>
                        <Typography textAlign="center">settings</Typography>
                    </MenuItem> */}
                    <MenuItem key="logout" onClick={logout}>
                        <Typography textAlign="center">logout</Typography>
                    </MenuItem>
                </Menu>
            </Box>
        </Box>
    );
}