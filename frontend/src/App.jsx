import React from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

import Home from './route/Home';
import Password from './route/Password';
import Settings from './route/Settings';
import TokenCapture from './route/token/Capture';
import TokenCreate from './route/token/Create';
import TokenExport from './route/token/Export';
import TokenImport from './route/token/Import';
import TokenInfo from './route/token/Info';
import TokenPassword from './route/token/Password';
import TokenRemove from './route/token/Remove';
import TokenUpdate from './route/token/Update';

import './App.css';

const App = () => {
  return (
    <Router>
        <div>
            <nav>
                <ol>
                    <li>
                        <Link to="/">Home</Link>
                    </li>
                    <li>
                        <Link to="/password">Password</Link>
                    </li>
                    <li>
                        <Link to="/token/capture">token capture</Link>
                    </li>
                    <li>
                        <Link to="/token/create">token create</Link>
                    </li>
                    <li>
                        <Link to="/token/info">token info</Link>
                    </li>
                    <li>
                        <Link to="/token/password">Token Password</Link>
                    </li>
                    <li>
                        <Link to="/token/update">Token Update</Link>
                    </li>
                    <li>
                        <Link to="/token/remove">Token Remove</Link>
                    </li>
                    <li>
                        <Link to="/token/import">Token Import</Link>
                    </li>
                    <li>
                        <Link to="/settings">Settings</Link>
                    </li>
                </ol>
            </nav>

            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/password" element={<Password />} />
                <Route path="/token/capture" element={<TokenCapture />} />
                <Route path="/token/create" element={<TokenCreate />} />
                <Route path="/token/info" element={<TokenInfo />} />
                <Route path="/token/password" element={<TokenPassword />} />
                <Route path="/token/update" element={<TokenUpdate />} />
                <Route path="/token/remove" element={<TokenRemove />} />
                <Route path="/token/export" element={<TokenExport />} />
                <Route path="/token/import" element={<TokenImport />} />
                <Route path="/settings" element={<Settings />} />
            </Routes>

        </div>
    </Router>
  );
};

export default App
