import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Navigate } from "react-router";
import { Route, Routes } from "react-router-dom";
import { useStore } from "./store/store";
import Login from "./pages/login/login.jsx";
import Signup from './pages/signup/signup';

function App() {

  const [state] = useStore();
  const {user: currentUser } = state;

  return (
    <React.Suspense fallback= {<div>Loading...</div>}>
      <Routes>
        {!currentUser ?
        <>
            <Route
              path="/login"
              element={<Login />}
            />
            <Route
              path="/signup"
              element={<Signup />}
            />
            <Route
              path="/"
              element={<Login />}
            />
            <Route
              path="*"
              element={<Login />}
            />
        </> :
        <>
            <Route
              path="/mainpage"
              element={<Login />}
            />
        </>
        }
      </Routes>
    </React.Suspense>
  )
}

export default App;
