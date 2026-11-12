const server: string = "http://localhost:8000"

type TokenResponse = {
    token: string;
    url: string
};

async function getToken(name: string,
                        room?: string): Promise<TokenResponse> {
    const res = await fetch(new URL("/api/token", server), {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({name, room}),
    });
    if (!res.ok) {
        const body = await res.json().catch(() => {});
        throw new Error(body.error ?? `HTTP ${res.status}`)
    }
    return res.json();
}