// DOM Elements
const submitButton = document.querySelector('#submit-button');
const heartDiseaseSelect = document.querySelector('input#heart-disease');
const fromInputs = document.querySelectorAll('input');

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
  const body = Array.from(fromInputs).reduce((acc, input) => {
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
      console.log("Successful POST", res)
    })
    .catch(err => {
      console.error(err);
    })
}
