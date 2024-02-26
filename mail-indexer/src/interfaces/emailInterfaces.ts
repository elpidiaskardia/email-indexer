interface EmailResult {
    total: number;
    data: any[];
}

export interface EmailData {
    status: string;
    result: EmailResult;
    message: string;
}
