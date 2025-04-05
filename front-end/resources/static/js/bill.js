function toggleModal(modalId, show) {
    document.getElementById(modalId).style.display = show ? 'block' : 'none';
}

function handleFormSubmission(formId, endpoint, data) {
    fetch(endpoint, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    }).then(response => {
        const success = response.ok;
        toggleModal(formId, false);
        document.getElementById(success ? 'successMessage' : 'errorMessage').style.display = 'block';
        setTimeout(() => {
            document.getElementById(success ? 'successMessage' : 'errorMessage').style.display = 'none';
        }, 1000);
    }).catch(() => {
        document.getElementById('errorMessage').style.display = 'block';
        setTimeout(() => {
            document.getElementById('errorMessage').style.display = 'none';
        }, 1000);
    });
}

document.getElementById('addMemberBtn').addEventListener('click', () => toggleModal('addMemberForm', true));
document.getElementById('cancelBtn').addEventListener('click', () => toggleModal('addMemberForm', false));
document.getElementById('addBillBtn').addEventListener('click', () => toggleModal('addBillForm', true));
document.getElementById('cancelBillBtn').addEventListener('click', () => toggleModal('addBillForm', false));

document.getElementById('memberForm').addEventListener('submit', function(event) {
    event.preventDefault();
    handleFormSubmission('addMemberForm', brokerEndpoint + '/members', {
        name: document.getElementById('name').value,
        email: document.getElementById('email').value
    });
});

document.getElementById('billForm').addEventListener('submit', function(event) {
    event.preventDefault();
    handleFormSubmission('addBillForm', brokerEndpoint + '/bills', {
        owner: document.getElementById('owner').value,
        members: Array.from(document.getElementById('members').selectedOptions).map(option => option.value),
        amount: document.getElementById('amount').value,
        notes: document.getElementById('notes').value
    });
});