package generatewithllm

const (
	webshop = `CREATE TABLE Products (
    product_id INT PRIMARY KEY,
    product_name VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    stock_quantity INT,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES Categories(category_id)
);

CREATE TABLE Categories (
    category_id INT PRIMARY KEY,
    category_name VARCHAR(255),
    parent_category_id INT
);

CREATE TABLE Customers (
    customer_id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    password VARCHAR(255), -- It's better to store hashed passwords
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20)
);

CREATE TABLE Orders (
    order_id INT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    status VARCHAR(50),
    total_amount DECIMAL(10, 2),
    shipping_address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20),
    payment_method VARCHAR(100),
    FOREIGN KEY (customer_id) REFERENCES Customers(customer_id)
);

CREATE TABLE Order_Details (
    order_detail_id INT PRIMARY KEY,
    order_id INT,
    product_id INT,
    quantity INT,
    unit_price DECIMAL(10, 2),
    subtotal DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (product_id) REFERENCES Products(product_id)
);`

	logistics = `CREATE TABLE Shipments (
    shipment_id INT PRIMARY KEY,
    shipment_date DATE,
    sender_id INT,
    recipient_id INT,
    delivery_status VARCHAR(50),
    shipping_method VARCHAR(100),
    total_weight DECIMAL(10, 2),
    shipping_cost DECIMAL(10, 2),
    FOREIGN KEY (sender_id) REFERENCES Customers(customer_id),
    FOREIGN KEY (recipient_id) REFERENCES Customers(customer_id)
);

CREATE TABLE Customers (
    customer_id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20)
);

CREATE TABLE Addresses (
    address_id INT PRIMARY KEY,
    customer_id INT,
    address_type VARCHAR(50),
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20),
    FOREIGN KEY (customer_id) REFERENCES Customers(customer_id)
);

CREATE TABLE Packages (
    package_id INT PRIMARY KEY,
    shipment_id INT,
    package_weight DECIMAL(10, 2),
    package_dimensions VARCHAR(100),
    package_contents TEXT,
    FOREIGN KEY (shipment_id) REFERENCES Shipments(shipment_id)
);

CREATE TABLE Carriers (
    carrier_id INT PRIMARY KEY,
    carrier_name VARCHAR(100),
    tracking_url VARCHAR(255),
    contact_information VARCHAR(255)
);

CREATE TABLE Shipment_Carriers (
    shipment_carrier_id INT PRIMARY KEY,
    shipment_id INT,
    carrier_id INT,
    tracking_number VARCHAR(100),
    FOREIGN KEY (shipment_id) REFERENCES Shipments(shipment_id),
    FOREIGN KEY (carrier_id) REFERENCES Carriers(carrier_id)
);
`

	hostpital = `-- Table for Patients
CREATE TABLE Patients (
    patient_id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE,
    gender VARCHAR(10),
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20)
);

-- Table for Doctors
CREATE TABLE Doctors (
    doctor_id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    specialization VARCHAR(100),
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20)
);

-- Table for Appointments
CREATE TABLE Appointments (
    appointment_id INT PRIMARY KEY,
    patient_id INT,
    doctor_id INT,
    appointment_date DATETIME,
    reason VARCHAR(255),
    status VARCHAR(50),
    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES Doctors(doctor_id)
);

-- Table for Medical Records
CREATE TABLE MedicalRecords (
    record_id INT PRIMARY KEY,
    patient_id INT,
    doctor_id INT,
    record_date DATE,
    diagnosis TEXT,
    treatment TEXT,
    prescription TEXT,
    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (doctor_id) REFERENCES Doctors(doctor_id)
);

-- Table for Ward
CREATE TABLE Ward (
    ward_id INT PRIMARY KEY,
    ward_name VARCHAR(100),
    ward_type VARCHAR(50),
    capacity INT
);

-- Table for Admission
CREATE TABLE Admissions (
    admission_id INT PRIMARY KEY,
    patient_id INT,
    ward_id INT,
    admission_date DATETIME,
    discharge_date DATETIME,
    status VARCHAR(50),
    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (ward_id) RE
`
	medical_record = `-- Table for Patients
CREATE TABLE Patients (
    patient_id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE,
    gender VARCHAR(10),
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20)
);

-- Table for Physicians
CREATE TABLE Physicians (
    physician_id INT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    specialization VARCHAR(100),
    address VARCHAR(255),
    city VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    phone_number VARCHAR(20)
);

-- Table for MedicalRecords
CREATE TABLE MedicalRecords (
    record_id INT PRIMARY KEY,
    patient_id INT,
    physician_id INT,
    record_date DATETIME,
    diagnosis TEXT,
    treatment TEXT,
    prescription TEXT,
    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (physician_id) REFERENCES Physicians(physician_id)
);

-- Table for LabTests
CREATE TABLE LabTests (
    test_id INT PRIMARY KEY,
    patient_id INT,
    physician_id INT,
    test_date DATETIME,
    test_name VARCHAR(255),
    result TEXT,
    FOREIGN KEY (patient_id) REFERENCES Patients(patient_id),
    FOREIGN KEY (physician_id
`
)

var KnownSchema = map[string]string{
	"webshop":        webshop,
	"logistics":      logistics,
	"hospital":       hostpital,
	"medical_record": medical_record,
}
