
import mongoose from "mongoose";

const { Schema } = mongoose;

    

export const addressSchema = new Schema(
  {
    country: String,
    state: String,
    pin: Number,    
  },
  {   
    strict: false
  }
);


export const addressModel = mongoose.model("address", addressSchema);
