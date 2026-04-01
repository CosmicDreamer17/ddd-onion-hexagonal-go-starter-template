"use client";

import { useState } from "react";
import type { RegisterRequest, UserResponse } from "../types/generated/index";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export default function RegisterForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [status, setStatus] = useState<"idle" | "loading" | "success" | "error">("idle");
  const [message, setMessage] = useState("");
  const [user, setUser] = useState<UserResponse | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setStatus("loading");
    setMessage("");
    setUser(null);

    const payload: RegisterRequest = { email, password };

    try {
      const res = await fetch(`${API_URL}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        const errText = await res.text();
        throw new Error(errText || "Registration failed");
      }

      const data: UserResponse = await res.json();
      setUser(data);
      setStatus("success");
      setMessage("Registration successful!");
      setEmail("");
      setPassword("");
    } catch (err: unknown) {
      setStatus("error");
      if (err instanceof Error) {
        setMessage(err.message);
      } else {
        setMessage("An unexpected error occurred");
      }
    }
  };

  return (
    <div className="w-full max-w-md mx-auto mt-10 p-6 bg-white rounded-lg shadow-md border border-gray-200">
      <h2 className="text-2xl font-bold mb-6 text-gray-800 text-center">Create an Account</h2>
      
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
          <input
            type="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-colors"
            placeholder="you@example.com"
          />
        </div>
        
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Password</label>
          <input
            type="password"
            required
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:border-blue-500 outline-none transition-colors"
            placeholder="••••••••"
          />
        </div>

        <button
          type="submit"
          disabled={status === "loading"}
          className="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-md transition-colors disabled:opacity-50"
        >
          {status === "loading" ? "Registering..." : "Register"}
        </button>
      </form>

      {status === "error" && (
        <div className="mt-4 p-3 bg-red-50 text-red-700 rounded-md text-sm border border-red-200">
          {message}
        </div>
      )}

      {status === "success" && user && (
        <div className="mt-4 p-4 bg-green-50 text-green-800 rounded-md text-sm border border-green-200">
          <p className="font-semibold">{message}</p>
          <div className="mt-2 text-xs font-mono break-all bg-green-100 p-2 rounded">
            ID: {user.id}<br/>
            Email: {user.email}<br/>
            Created: {new Date(user.created_at).toLocaleString()}
          </div>
        </div>
      )}
    </div>
  );
}
