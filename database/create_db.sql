CREATE TABLE Users (
    ID SERIAL PRIMARY KEY,
    Telegram_Username VARCHAR(255),
    Nickname VARCHAR(255),
    UserType VARCHAR(255)
);

CREATE TABLE Cards (
    ID SERIAL PRIMARY KEY,
    Card_Number VARCHAR(255),
    Issuing_Bank VARCHAR(255),
    Daily_Limit NUMERIC,
    Current_Balance NUMERIC,
    Daimyo_ID INTEGER REFERENCES Users(ID)
);

CREATE TABLE Sessions (
    ID SERIAL PRIMARY KEY,
    User_ID INTEGER REFERENCES Users(ID),
    Entity_Type VARCHAR(255),
    Start_Date_Time TIMESTAMP,
    End_Date_Time TIMESTAMP
);

CREATE TABLE Actions (
    ID SERIAL PRIMARY KEY,
    Session_ID INTEGER REFERENCES Sessions(ID),
    Entity_ID INTEGER REFERENCES Users(ID),
    Action_Type VARCHAR(255),
    Action_Date_Time TIMESTAMP
);

CREATE TABLE Relations (
    ID SERIAL PRIMARY KEY,
    Entity_ID INTEGER REFERENCES Users(ID),
    Entity_Nickname varchar(255),
    Related_Entity_ID INTEGER REFERENCES Users(ID),
    Related_Entity_Nickname varchar(255),
    Relation_Type VARCHAR(255),
    Creation_Date_Time TIMESTAMP
);