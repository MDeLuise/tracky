import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Container from '@mui/material/Container';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import { AxiosInstance } from 'axios';
import { NavigateFunction, useNavigate } from 'react-router-dom';
import secureLocalStorage from 'react-secure-storage';


export default function Navbar(props: { requestor: AxiosInstance }) {
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
    }

    const getGravatarUrl = (): string => {
        let md5 = require('md5');
        return `https://www.gravatar.com/avatar/${md5(username)}?d=identicon`;
    }

    return (
        <AppBar position="static">
            <Container maxWidth="xl">
                <Toolbar disableGutters>
                    <Typography
                        variant="h6"
                        noWrap
                        component="a"
                        href="/"
                        sx={{
                            mr: 2,
                            display: { xs: 'none', md: 'flex' },
                            fontWeight: 700,
                            letterSpacing: '.3rem',
                            color: 'inherit',
                            textDecoration: 'none',
                        }}
                    >
                        Tracky
                    </Typography>

                    <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
                        <IconButton
                            size="large"
                            aria-label="account of current user"
                            aria-controls="menu-appbar"
                            aria-haspopup="true"
                            onClick={handleOpenNavMenu}
                            color="inherit"
                        >
                            <MenuIcon />
                        </IconButton>
                        <Menu
                            id="menu-appbar"
                            anchorEl={anchorElNav}
                            anchorOrigin={{
                                vertical: 'bottom',
                                horizontal: 'left',
                            }}
                            keepMounted
                            transformOrigin={{
                                vertical: 'top',
                                horizontal: 'left',
                            }}
                            open={Boolean(anchorElNav)}
                            onClose={handleCloseNavMenu}
                            sx={{
                                display: { xs: 'block', md: 'none' },
                            }}
                        >
                            <MenuItem key="trackers" onClick={() => {
                                handleCloseNavMenu();
                                navigate("/trackers");
                            }}>
                                <Typography textAlign="center">All Trackers</Typography>
                            </MenuItem>
                            <MenuItem key="observations" onClick={() => {
                                handleCloseNavMenu();
                                navigate("/observations");
                            }}>
                                <Typography textAlign="center">All Observations</Typography>
                            </MenuItem>
                        </Menu>
                    </Box>
                    <Typography
                        variant="h5"
                        noWrap
                        component="a"
                        href="/"
                        sx={{
                            mr: 2,
                            display: { xs: 'flex', md: 'none' },
                            flexGrow: 1,
                            fontWeight: 700,
                            letterSpacing: '.3rem',
                            color: 'inherit',
                            textDecoration: 'none',
                        }}
                    >
                        Tracky
                    </Typography>
                    <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
                        <Button
                            key="trackers"
                            onClick={() => {
                                handleCloseNavMenu();
                                navigate("/trackers");
                            }}
                            sx={{ my: 2, color: 'white', display: 'block' }}
                        >
                            Trackers
                        </Button>
                        <Button
                            key="observations"
                            onClick={() => {
                                handleCloseNavMenu();
                                navigate("/observations");
                            }}
                            sx={{ my: 2, color: 'white', display: 'block' }}
                        >
                            Observations
                        </Button>
                    </Box>

                    <Box sx={{ flexGrow: 0 }}>
                        <Tooltip title="Open settings">
                            <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                                <Avatar alt={username} src={getGravatarUrl()} />
                            </IconButton>
                        </Tooltip>
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
                            {/* <MenuItem key="settings" onClick={handleCloseUserMenu}>
                                <Typography textAlign="center">settings</Typography>
                            </MenuItem> */}
                            <MenuItem key="logout" onClick={logout}>
                                <Typography textAlign="center">logout</Typography>
                            </MenuItem>
                        </Menu>
                    </Box>
                </Toolbar>
            </Container>
        </AppBar>
    );
}