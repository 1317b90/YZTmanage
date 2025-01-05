export interface taskI {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    Type: string;
    Input: string;
    Output: string;
    State: string;
}

export interface filterI {
    text: string;
    value: string;
}



export interface userI {
    Phone: string;
    Enable: boolean;
    Uscid: string;
    DsjUsername: string;
    DsjPassword: string;
    CompanyName: string;
    BankName: string;
    BankID: string;
}



