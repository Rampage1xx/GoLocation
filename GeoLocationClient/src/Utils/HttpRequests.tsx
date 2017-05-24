import axios from 'axios';

// creating a http instance
 const instance1 = axios.create({
    baseURL: 'http://localhost:9000',
    headers: {'Content-Type': 'application/json'},
    validateStatus: (status) => {
        return status >= 200 && status < 300;
    },
    withCredentials: true

});

// fetching the markers from the backend

 const getMarkers = () => {
    return instance1.get('/getData')
        .then(r => r)
        .catch(e => e);
};

 export {instance1, getMarkers}
