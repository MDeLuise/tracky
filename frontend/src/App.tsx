import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import "./style/Base.scss";
import { BrowserRouter, Routes, Route, NavigateFunction, useNavigate } from "react-router-dom";
import Auth from "./components/Auth";
import Home from "./components/Home";
import axios from "axios";
import Settings from "./components/Settings";
import secureLocalStorage from "react-secure-storage";
import React from "react";
import { ThemeProvider, createTheme } from '@mui/material/styles';
import Tracker from "./components/Tracker";
import AllTrackers from "./components/AllTrackers";
import {CssBaseline } from "@mui/material";
import AllObservations from "./components/AllObservations";
import AddEntities from "./components/AddEntities";

const ColorModeContext = React.createContext({ toggleColorMode: () => { } });

export function App() {
  const isLoggedIn: () => boolean = () => secureLocalStorage.getItem("tracky-key") != null;
  const backendURL = window._env_.REACT_APP_API_URL != null ? window._env_.REACT_APP_API_URL : "http://localhost:8085/api";
  const axiosReq = axios.create({
    baseURL: backendURL,
    timeout: 1000
  });

  axiosReq.interceptors.request.use(
    (req) => {
      if (!req.url?.startsWith("authentication") && !req.url?.startsWith("api-key")) {
        req.headers['Key'] = secureLocalStorage.getItem("tracky-key");
      }
      return req;
    },
    (err) => {
      return Promise.reject(err);
    }
  );

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home requestor={axiosReq} isLoggedIn={isLoggedIn} />} />
        <Route path="/tracker/:trackerId" element={<Tracker requestor={axiosReq} isLoggedIn={isLoggedIn} />} />
        <Route path="/trackers" element={<AllTrackers requestor={axiosReq} isLoggedIn={isLoggedIn} />} />
        <Route path="/observations" element={<AllObservations requestor={axiosReq} isLoggedIn={isLoggedIn} />} />
        <Route path="/auth" element={<Auth requestor={axiosReq} />} />
        <Route path="/settings" element={<Settings requestor={axiosReq} isLoggedIn={isLoggedIn} />} />
        <Route path="/add" element={<AddEntities requestor={axiosReq} isLoggedIn={isLoggedIn} />} />
      </Routes>
    </BrowserRouter>
  );
}


export default function AppWithColorMode() {
  const [mode, setMode] = React.useState<'light' | 'dark'>(localStorage.getItem("tracky-dark") != "false" ? "dark" : "light");
  const colorMode = React.useMemo(
    () => ({
      toggleColorMode: () => {
        setMode((prevMode) => (prevMode === 'light' ? 'dark' : 'light'));
      },
    }),
    [],
  );

  const theme = React.useMemo(
    () =>
      createTheme({
        palette: {
          mode,
          ...(mode === 'light'
            ? {
              // palette values for light mode
              primary: {
                main: '#3f51b5',
              },
              secondary: {
                main: '#f50057',
              },
              background: {
                default: '#303030',
              }
            }
            : {
              // palette values for dark mode
              primary: {
                main: '#3f51b5',
              },
              secondary: {
                main: '#f50057',
              },
              background: {
                default: "#303030",
                paper: "#424242",
              }
            }),
        }
      }),
    [mode],
  );

  return (
    <ColorModeContext.Provider value={colorMode}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <App></App>
      </ThemeProvider>
    </ColorModeContext.Provider>
  );
}
