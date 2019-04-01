// DOM Elements
const submitButton = document.querySelector('#submit-button');
const heartDiseaseSelect = document.querySelector('input#heart-disease');
const formInputs = Array.from(document.querySelectorAll('input'));
const predictionSection = document.querySelector('#prediction-results');
const predictionVal = document.querySelector('#prediction-value');
const scoreVal = document.querySelector('#score-value');

// Event Listeners
heartDiseaseSelect.addEventListener('input', selectHeartDiseaseStatus);

// Event Handlers
function selectHeartDiseaseStatus(event) {
  if (event.target.value) {
    submitButton.innerText = "ENTER NEW DATA POINT"
  } else {
    submitButton.innerText = "MAKE PREDICTION"
  }
}

function sendData() {
  const body = formInputs.reduce((acc, input) => {
    acc[input.name] = input.value;
    return acc;
  }, {});

  fetch('http://localhost:8000/patients', {
    method: 'POST',
    body: JSON.stringify(body),
    headers: {
      "Content-Type": "text/csv",
    }
  })
    .then(res => {
      clearInputs();
      return res.json();
    })
    .then((data) => {
      const { prediction, score } = data;
      predictionVal.innerText = prediction;
      scoreVal.innerText = score;
      togglePredictionSection();
    })
    .catch(err => {
      console.error(err);
    })
}

function clearInputs() {
  formInputs.forEach(input => input.value = '');
}

function togglePredictionSection() {
  if (!predictionVal.innerText || !scoreVal.innerText) {
    predictionSection.style.display = 'none';
  } else {
    predictionSection.style.display = 'block';
  }
}

togglePredictionSection();