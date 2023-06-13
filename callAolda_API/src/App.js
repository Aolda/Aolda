import React, { useState } from 'react';
import logo from './images/aolda_logo.png';
import scroll from './images/scroll.png';
import './App.css';

function Tooltip({ children, tooltip }) {
  const [showTooltip, setShowTooltip] = useState(false);

  const handleMouseEnter = () => {
    setShowTooltip(true);
  };

  const handleMouseLeave = () => {
    setShowTooltip(false);
  };

  const renderTooltipContent = () => {
    const lines = tooltip.split('\n'); // 입력된 값을 줄바꿈을 기준으로 나눔

    return lines.map((line, index) => (
      <React.Fragment key={index}>
        {line}
        <br /> {/* 줄바꿈 표시 */}
      </React.Fragment>
    ));
  };

  return (
    <div className="tooltip-container">
      <div
        className="tooltip-trigger"
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
      >
        {children}
      </div>
      {showTooltip && <div className="tooltip">{renderTooltipContent()}</div>}
    </div>
  );
}


function App() {
  const [fileHash, setFileHash] = useState('');
  const [functionName, setFunctionName] = useState('');
  const [args, setArgs] = useState([]);
  const [result, setResult] = useState('');
  const [activeTab, setActiveTab] = useState(0);
  const [selectedFile, setSelectedFile] = useState(null);
  const [fileTx, setFileTx] = useState('');
  const [resultTx, setResultTx] = useState('');

  
  const jsonData = {
    fileHash: fileHash,
    functionName: functionName,
    args: args
  };

  const handleTabChange = (index) => {
    setActiveTab(index);
  };

  const handleFileHash = (event) => {
    setFileHash(event.target.value);
  };

  const handleFunctionName = (event) => {
    setFunctionName(event.target.value);
  };

  // const handleArgs = (event) => {
  //   setArgs(event.target.value);
  // };
  const handleArgs = (event) => {
    setArgs(event.target.value.split('\n')); // 줄바꿈을 기준으로 문자열을 배열로 분할하여 설정
  };
  
  const handleScrollDown = () => {
    const targetElement = document.getElementById('body');
    targetElement.scrollIntoView({ behavior: 'smooth' });
  };

  const handleScrollUp = () => {
    const targetElement = document.getElementById('header');
    targetElement.scrollIntoView({ behavior: 'smooth' });
  };

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    setSelectedFile(file);
  };

  const handleUploadButton = () => {
    const formData = new FormData();
    formData.append('file', selectedFile);

    fetch('http://localhost:8000/upload', {
      method: 'POST',
      body: formData
    })
      .then(response => response.text())
      .then(data => {
        console.log('POST 요청 결과:', data);
        setFileTx(data);
      })
      .catch(error => {
        console.error('POST 요청 중 오류 발생:', error);
      });
  };

  const handleButton = () => {
    console.log(jsonData);
    fetch('http://localhost:8000/emit', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(jsonData)
    })
      .then(response => response.text())
      .then(data => {
        console.log('POST 요청 결과:', data);
        setResult(data);
      })
      .catch(error => {
        console.error('POST 요청 중 오류 발생:', error);
      });
  };

  const parsedData = fileTx ? JSON.parse(fileTx) : null;

  return (
    <div className="App">
      <header className="App-header" id="header">
        <p className="font-title">{"Aolda-FlexiContract"}</p>
        <p className="font-subtitle">
          {"Solidity 외의 언어로 작성된 Smart Contract를 EVM 기반 체인에 적용하기 위한 L2 솔루션"}
        </p>
        <img src={logo} className="App-logo" alt="logo" />
        <p className="font-link">
          아올다팀의 FlexiContract을 소개하는 링크입니다 :)
        </p>
        <a
          className="App-link"
          href="https://github.com/Aolda"
          target="_blank"
          rel="noopener noreferrer"
        >
          GitHub
        </a>
        <a
          className="App-link"
          href="https://velog.io/@aolda_aww"
          target="_blank"
          rel="noopener noreferrer"
        >
          Velog
        </a>
        <div onClick={handleScrollDown} className="scrolldown-button">
          <img src={scroll} alt="scroll" className="scroll-img" />
        </div>
      </header>

      <div className="App-container" id="body">
        <div onClick={handleScrollUp} className="scrollup-button">
          <img src={scroll} alt="scroll" className="scroll-img" />
        </div>

        <div className="tabs">
          <p
            style={{ color: activeTab === 0 ? '#61dafb' : 'black' }}
            className="tab"
            onClick={() => handleTabChange(0)}
          >
            파일 등록
          </p>
          <p
            style={{ color: activeTab === 1 ? '#61dafb' : 'black' }}
            className="tab"
            onClick={() => handleTabChange(1)}
          >
            callAolda
          </p>
        </div>
        <div>
          {activeTab === 0 && (
            <div className="tab-content">
              <p className="font-1">
                FlexiContract에 업로드 할 파일을 추가해주세요!
              </p>

              <input type="file" onChange={handleFileChange} />
              {selectedFile && <p>선택한 파일: {selectedFile.name}</p>}

              <button className="button-emit" onClick={handleUploadButton}>
                upload File
              </button>

              <div className={`transaction-section ${fileTx ? 'visible' : ''}`}>
                <p className="font-inputtitle">Transaction</p>
                {parsedData && (
                  <div>
                    <Tooltip tooltip={"Type은 트랜잭션의 종류를 나타내는 값으로 0~4 중 하나입니다.\n 0: 파일 생성\n 1: EVM에서 AOLDA를 호출한 기록\n 2: USER가 API를 사용해 직접 Aolda Node를 호출한 기록\n 3: type1과 type2에 대한 결과값\n 4: 블록 채굴에 대한 트랜잭션"}><strong>Type: </strong>{parsedData.header.type}</Tooltip>
                    <Tooltip tooltip={"Hash는 트랜잭션 body에 대한 해쉬값입니다."}><strong>Hash:</strong> {parsedData.header.hash}</Tooltip>
                    <Tooltip tooltip={"Block Number는 local blockchain의 block 개수입니다."}><strong>Block Number: </strong>{parsedData.header.blockNumber}</Tooltip>
                    <Tooltip tooltip={"Transaction Index은 해당 Transaction의 순서입니다"}><strong>Transaction Index: </strong>{parsedData.header.transactionIndex}</Tooltip>
                    <Tooltip tooltip={"From은 트랜잭션을 생성한 Node입니다."}><strong>From: </strong>{parsedData.header.from}</Tooltip>
                    <Tooltip tooltip={"Nonce는 해당 트랜잭션을 채굴할 때 쓰이는 해쉬값과 연관된 값입니다."}><strong>Nonce: </strong>{parsedData.header.nonce}</Tooltip>
                    <Tooltip tooltip={"Signature은 트랜잭션 유효성 검증을 위한 전자 서명입니다."}><strong>Signature R: </strong>{parsedData.header.signature.R}</Tooltip>
                    <Tooltip tooltip={"Signature은 트랜잭션 유효성 검증을 위한 전자 서명입니다."}><strong>Signature S: </strong>{parsedData.header.signature.S}</Tooltip>
                    <Tooltip tooltip={"Signature은 트랜잭션 유효성 검증을 위한 전자 서명입니다."}><strong>Signature V: </strong>{parsedData.header.signature.V}</Tooltip>
                    <Tooltip tooltip={"Timestamp는 트랜잭션이 생성된 시간입니다."}><strong>Timestamp: </strong>{parsedData.header.timeStamp}</Tooltip>
                    <Tooltip tooltip={"File Hash는 파일 공유 분산 네트워크인 IPFS에서 반환된 fileHash입니다.\ntype0 : 파일을 IPFS 네트워크에 저장하고 반환된 Hash값입니다.\ntype1, type2, type3 : 호출할 때 입력된 인자 중 fileHash에 해당하는 값입니다.\ntype4 : null값입니다."}><strong>File Hash: </strong>{parsedData.body.fileHash}</Tooltip>
                    <Tooltip tooltip={"Function Name는 인자값으로 type1, type2, type3에서만 사용됩니다.\ntype0, type4 : null값입니다.\ntype1, type2, type3 : 호출할 때 입력된 인자 중 functionName에 해당하는 값입니다."}><strong>Function Name: </strong>{parsedData.body.functionName}</Tooltip>
                    <Tooltip tooltip={"Arguments는 인자값으로 type1, type2, type3에서만 사용됩니다.\ntype0, type4 : null값입니다.\ntype1, type2, type3 : 호출할 때 입력된 인자 중 arguments에 해당하는 값입니다. \n인자들은 모두 string으로 들어와야 합니다."}><strong>Arguments: </strong>{parsedData.body.arguments}</Tooltip>
                    <Tooltip tooltip={"Result는 결과값입니다.\ntype0, type1, type2 : null값입니다.\ntype3 : 함수 호출에 대한 결과값입니다.\ntype4 : 보상으로 지급되는 토큰의 수입니다."}><strong>Result: </strong>{parsedData.body.result}</Tooltip>
                  </div>
                )}
              </div>
            </div>
          )}

          {activeTab === 1 && (
            <div className="tab-content">
              <p className="font-1">필요한 데이터를 입력해주세요!</p>

              <div>
                <div className="container-input">
                  <span className="font-inputtitle">fileHash :</span>
                  <input
                    className="input"
                    type="text"
                    value={fileHash}
                    onChange={handleFileHash}
                  />
                </div>
                <div className="container-input">
                  <span className="font-inputtitle">functionName :</span>
                  <input
                    className="input"
                    type="text"
                    value={functionName}
                    onChange={handleFunctionName}
                  />
                </div>
                <div className="container-input">
                  <span className="font-inputtitle">args :</span>
                  <input
                    className="input"
                    type="text"
                    value={args}
                    onChange={handleArgs}
                  />
                </div>
              </div>

              <button className="button-emit" onClick={handleButton}>
                callAolda
              </button>

              <div>
                <p className="font-inputtitle">
                  Aolda를 통해 나온 결과 : {result}
                </p>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
