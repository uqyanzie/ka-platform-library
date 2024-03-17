import React from 'react';
import './App.css';

import { PlatformListRequest, PlatformListResponse } from './pb/platform_pb';

// import { proto } from './pb/platform';

import { PlatformServiceClient } from './pb/PlatformServiceClientPb';

const client = new PlatformServiceClient('http://localhost:8080');

function getPlatformList() {
  return new Promise<PlatformListResponse>((resolve, reject) => {
    client.getPlatformList(new PlatformListRequest(), {}, (err, response) => {
      if (err) {
        reject(err);
      } else {
        resolve(response);
      }
    });
  });
}

function App() {

  const [platforms, setPlatforms] = React.useState<PlatformListResponse | null>(null);

  React.useEffect(() => {
    console.log(client.methodDescriptorGetPlatformList.getName());
    getPlatformList().then((response) => {
      setPlatforms(response);
    });
  }, []);

  return (
    <div className="App">
      <h1 className='text-center'>Platform List</h1>
      <table style={{marginBottom: 20}}>
        <thead>
          <tr>
            <th>Unit Name</th>
            <th>Unit Class</th>
            <th>Operation Field</th>
            <th>General Category</th>
            <th>General Type</th>
            <th>Lethality Level</th>
            <th>Cruising Speed (Knots)</th>
            <th>Maximum Speed (Knots)</th>
            <th>Combat Range (Nautical Miles)</th>
            <th>Fuel Load (liters)</th>
            <th>Country Origin</th>
          </tr>
        </thead>
        <tbody>
          {platforms && platforms.getPlatformsList().map((platform) => (
            <tr key={platform.getId()}>
              <td>{platform.getUnitname()}</td>
              <td>{platform.getUnitclass()}</td>
              <td>{platform.getOperationField()}</td>
              <td>{platform.getGeneralcategory()}</td>
              <td>{platform.getGeneraltype()}</td>
              <td>{platform.getLethalitylevel()}</td>
              <td>{platform.getCruisingspeed()}</td>
              <td>{platform.getMaxspeed()}</td>
              <td>{platform.getCombatrange()}</td>
              <td>{platform.getFuelload()}</td>
              <td>{platform.getCountryorigin()}</td>
            </tr>
          ))}
        </tbody>
      </table>
      
      <div style={{display:"flex"}}>
        <button style={{marginLeft:"auto", marginRight:"auto", fontSize:"20px"}}>
          + Add New Platform
        </button>
      </div>
      
    </div>
  );
}

export default App;
