import React from "react";

export default function Layout({ title, children }: { title: string, children: React.ReactNode }) {
  return (
    <div className="min-h-screen flex justify-center">
      <div className="mx-auto space-y-3 pt-[130px]">
        <h1 className="text-black font-medium text-[60px] text-left">{title}</h1>
        {children}
      </div>
    </div>
  );
}
