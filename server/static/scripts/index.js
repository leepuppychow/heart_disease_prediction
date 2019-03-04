// DOM Elements
const submitButton = document.querySelector('#submit-button');
const heartDiseaseSelect = document.querySelector('select#heart-disease');

// Event Listeners
heartDiseaseSelect.addEventListener('change', selectHeartDiseaseStatus);

// Event Handlers
function selectHeartDiseaseStatus(event) {
  if (event.target.value) {
    submitButton.value = "ENTER NEW DATA POINT"
  } else {
    submitButton.value = "MAKE PREDICTION"
  }
}
