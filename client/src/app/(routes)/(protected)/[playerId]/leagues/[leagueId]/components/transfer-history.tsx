"use client";

import React from "react";

const TransferHistorySection: React.FC = () => {
  return (
    <section className="w-full">
      <h2 className="text-xl font-semibold mb-4">Transfer History</h2>
      <div className="overflow-x-auto">
        <p className="text-gray-500">No transfer history available</p>
      </div>
    </section>
  );
};

export default TransferHistorySection;
